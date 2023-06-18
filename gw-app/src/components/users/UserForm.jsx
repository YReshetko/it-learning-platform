import React, {useRef} from 'react';
import CreateUserTable from "./CreateUserTable";
import Button from "../../ui/button/Button";
import '../../utils/cookies'
import * as cookies from '../../utils/cookies';

const UserForm = () => {
    const tableCallback = useRef(null);
    const table = <CreateUserTable ref={tableCallback}/>

    const send = (value) => {
        console.log("send")
        let users = tableCallback.current.getData()
        console.log({users: users});

        let newUsers = users.map((user) => {
            return {
                login: user.login,
                first_name: user.firstName,
                last_name: user.lastName,
                email: user.email,
                roles: user.roles
            }
        })

        let request = {
            method: 'POST', // or GET
            redirect: 'follow', // React to redirect if 3xx status is returned
            body: JSON.stringify({users: newUsers}), // data can be string or object
            headers: {
                'Content-Type': 'application/json'
            }
        }
        // let accessToken = cookies.getCookie('access_token')
        let accessToken = window.localStorage.getItem("access_token")
        if (accessToken) {
            request.headers['Authorization'] = 'Bearer '+ accessToken;
        }

        fetch("/api/auth/users", request)
            .then(res => {
                res.json();
                if (res.redirected) {
                    window.location.replace(res.url);
                    //window.location.href = res.url;
                }
            }) // if response is json, for text use res.text()
            .then(response => console.log('Response:', JSON.stringify(response))) // if text, no need for JSON.stringify
            .catch(error => console.error('Error:', error));
    }

    const add = (value) => {
        tableCallback.current.addRow();

    }

    const remove = (value) => {
        tableCallback.current.removeRow()
    }
    return (
        <div>
            <h2>Создать пользователей</h2>
            {table}
            <Button onClick={send}>Отправить</Button>
            <Button onClick={add}>Добавить</Button>
            <Button onClick={remove}>Удалить</Button>
            <h3>Роли записываются через запятую: ADMIN MANAGER TEACHER STUDENT</h3>
        </div>
    );
};

export default UserForm;
