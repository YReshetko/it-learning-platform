import React, {useRef} from 'react';
import CreateUserTable from "./CreateUserTable";
import Button from "../../ui/button/Button";
import * as rest from '../../utils/rest';

const UserForm = () => {
    const tableCallback = useRef(null);
    const table = <CreateUserTable ref={tableCallback}/>

    const send = (value) => {
        let users = tableCallback.current.getData()
        let newUsers = users.map((user) => {
            return {
                login: user.login,
                first_name: user.firstName,
                last_name: user.lastName,
                email: user.email,
                roles: user.roles
            }
        })

        rest.post("/api/registration/users", {users: newUsers})
            .then(response => {
                console.log('Response:', JSON.stringify(response));
                tableCallback.current.clean();
            })
            .catch(error => {
                console.error('Error:', error)
            });
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
