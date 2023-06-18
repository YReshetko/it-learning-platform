import React from 'react';
import Button from "../../ui/button/Button";
import Label from "../../ui/label/Label";
import * as auth from "../../utils/auth.js"

const Header = ({firstName, lastName}) => {
    const logout = (event) => {
        auth.clean();
    }
    return (
        <div style={{display: "inline-block"}}>
            <Button onClick={logout}>Выйти</Button>
            <Label>{firstName} {lastName}</Label>
        </div>
    );
};

export default Header;