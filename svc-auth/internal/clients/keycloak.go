package clients

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/model"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

/*
KeycloakClient the client to keycloak service
@Constructor
*/
type KeycloakClient struct {
	cfg         config.KeycloakClient
	client      *gocloak.GoCloak // @Exclude
	clientToken token            // @Exclude
}

type token struct {
	clientToken *gocloak.JWT
	createdAt   time.Time
}

// @PostConstruct
func (kc *KeycloakClient) postConstruct() {
	kc.client = gocloak.NewClient(fmt.Sprintf("http://%s:%d", kc.cfg.Host, kc.cfg.Port))
}

func (kc *KeycloakClient) CreateUser(ctx context.Context, user model.User) (string, error) {
	clientToken, err := kc.getClientToken(ctx)
	if err != nil {
		return "", fmt.Errorf("unable create user: %w", err)
	}
	roles := user.Roles.ToStringPtr()

	keycloakUserID, err := kc.client.CreateUser(ctx, clientToken.AccessToken, kc.cfg.Realm, gocloak.User{
		Username:               toPtr(user.Login),
		Enabled:                toPtr(true),
		EmailVerified:          toPtr(true),
		FirstName:              toPtr(user.FirstName),
		LastName:               toPtr(user.LastName),
		Email:                  toPtr(user.Email),
		ServiceAccountClientID: toPtr(kc.cfg.ClientID),
	})

	if err != nil {
		return keycloakUserID, fmt.Errorf("unable to create user on keycloak: %w", err)
	}

	// TODO remove the password logging
	pass := generatePassword(6)
	err = kc.client.SetPassword(ctx, clientToken.AccessToken, keycloakUserID, kc.cfg.Realm, pass, true)
	if err != nil {
		return keycloakUserID, fmt.Errorf("unable to set initial password: %w", err)
	}
	fmt.Printf("User created: %s with password: %s\n", user.Login, pass) // TODO !!!!!!!!!!!!!!!!!!!!!

	err = kc.removeAllRealmRoles(ctx, keycloakUserID)
	if err != nil {
		return keycloakUserID, fmt.Errorf("unable to finalize keycloak user creation: %w", err)
	}

	userRoles, err := kc.findRealmUserRoles(ctx, roles)
	if err != nil {
		return keycloakUserID, fmt.Errorf("unable to finalize user roles: %w", err)
	}

	if len(userRoles) == 0 {
		return keycloakUserID, nil
	}

	err = kc.client.AddRealmRoleToUser(ctx, clientToken.AccessToken, kc.cfg.Realm, keycloakUserID, userRoles)
	if err != nil {
		return keycloakUserID, fmt.Errorf("unable to update user roles after creation: %w", err)
	}

	return keycloakUserID, nil
}

func (kc *KeycloakClient) ValidateAccessToken(ctx context.Context, accessToken string) (bool, error) {
	result, err := kc.client.RetrospectToken(ctx, accessToken, kc.cfg.ClientID, kc.cfg.ClientSecret, kc.cfg.Realm)
	if err != nil {
		return false, fmt.Errorf("unable to retrospect token: %w", err)
	}
	if result == nil {
		return false, errors.New("retrospect result is nil")
	}
	if !*result.Active {
		return false, nil
	}
	return true, nil
}

func (kc *KeycloakClient) GetUserIDAndRoles(ctx context.Context, accessToken string) (uuid.UUID, model.Roles, error) {
	userInfo, err := kc.client.GetUserInfo(ctx, accessToken, kc.cfg.Realm)
	if err != nil {
		return uuid.UUID{}, nil, fmt.Errorf("unable to get user info: %w", err)
	}
	if userInfo == nil {
		return uuid.UUID{}, nil, errors.New("no userInfo is returned")
	}
	userId, err := uuid.Parse(*userInfo.Sub)
	if err != nil {
		return uuid.UUID{}, nil, fmt.Errorf("unable to parse user ID from *userInfo.Sub: %w", err)
	}

	clientToken, err := kc.getClientToken(ctx)
	if err != nil {
		return uuid.UUID{}, nil, fmt.Errorf("unable to get client token: %w", err)
	}

	roles, err := kc.client.GetRealmRolesByUserID(ctx, clientToken.AccessToken, kc.cfg.Realm, userId.String())
	if err != nil {
		return userId, nil, fmt.Errorf("unable to get user realm roles: %w", err)
	}
	var outRoles model.Roles
	for _, modelRole := range model.AllRoles {
		for _, realmRole := range roles {
			if realmRole.Name != nil && *realmRole.Name == string(modelRole) {
				outRoles = append(outRoles, modelRole)
			}
		}
	}

	return userId, outRoles, nil
}

func (kc *KeycloakClient) removeAllRealmRoles(ctx context.Context, userID string) error {
	clientToken, err := kc.getClientToken(ctx)
	if err != nil {
		return fmt.Errorf("unable to remove user roles: %w", err)
	}

	roles, err := kc.client.GetRealmRolesByUserID(ctx, clientToken.AccessToken, kc.cfg.Realm, userID)
	if err != nil {
		return fmt.Errorf("unable to get user realm roles: %w", err)
	}

	rolesToDelete := make([]gocloak.Role, len(roles))
	for i, role := range roles {
		if role != nil {
			rolesToDelete[i] = *role
		}
	}

	err = kc.client.DeleteRealmRoleFromUser(ctx, clientToken.AccessToken, kc.cfg.Realm, userID, rolesToDelete)
	if err != nil {
		return fmt.Errorf("unable to remove user realm roles: %w", err)
	}
	return nil
}

func (kc *KeycloakClient) findRealmUserRoles(ctx context.Context, roles []string) ([]gocloak.Role, error) {
	if len(roles) == 0 {
		return nil, nil
	}
	clientToken, err := kc.getClientToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to find realm roles: %w", err)
	}

	realmRoles, err := kc.client.GetRealmRoles(ctx, clientToken.AccessToken, kc.cfg.Realm, gocloak.GetRoleParams{})
	if err != nil {
		return nil, fmt.Errorf("unable to get realm roles: %w", err)
	}

	var out []gocloak.Role
	for _, role := range realmRoles {
		if role == nil {
			continue
		}
		realmRole := *role
		for _, roleStr := range roles {
			if realmRole.Name != nil && *realmRole.Name == roleStr {
				out = append(out, realmRole)
			}
		}
	}
	return out, nil
}

func (kc *KeycloakClient) getClientToken(ctx context.Context) (*gocloak.JWT, error) {
	if kc.clientToken.isExpired() {
		t, err := kc.newClientToken(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to refresh client token: %w", err)
		}
		kc.clientToken = t
	}
	return kc.clientToken.clientToken, nil
}

func (kc *KeycloakClient) newClientToken(ctx context.Context) (token, error) {
	createdAt := time.Now()
	t, err := kc.client.GetToken(ctx, kc.cfg.Realm, gocloak.TokenOptions{
		ClientID:     &kc.cfg.ClientID,
		ClientSecret: &kc.cfg.ClientSecret,
		GrantType:    toPtr("client_credentials"),
		//Scope:         toPtr("profile"),
		//ResponseTypes: &[]string{"token", "id_token"},
	})
	if err != nil {
		return token{}, fmt.Errorf("unable to create new client access token: %w", err)
	}
	return token{
		clientToken: t,
		createdAt:   createdAt,
	}, nil
}

func toPtr[T any](v T) *T {
	return &v
}

func (t token) isExpired() bool {
	if t.clientToken == nil {
		return true
	}

	expireTime := t.createdAt.Add(time.Second * time.Duration(t.clientToken.ExpiresIn))
	if time.Now().After(expireTime) {
		return true
	}

	return false
}

func generatePassword(num int) string {
	symbols := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"01234567890",
		"_-!@#$%^&*",
	}
	password := ""
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < num; i++ {
		lineIndex := rnd.Intn(len(symbols))
		chIndex := rnd.Intn(len(symbols[lineIndex]))
		ch := symbols[lineIndex][chIndex]
		password += string(ch)
	}
	return password
}
