import React, {forwardRef, useImperativeHandle, useState} from 'react';
import CreateUserTableRow from "./CreateUserTableRow";

const CreateUserTable = forwardRef((props, ref) => {
    const [users, setUsers] = useState([{id: 1, login: '', firstName: '', lastName: '', email: '', roles: []}]);
    const updateUser = (id, field, value) => {
        console.log(id)
        console.log(field)
        console.log(value)
        users[id - 1][field] = value
        setUsers(users);
    }

    useImperativeHandle(ref, () => ({
        addRow() {
            let user = {id: users.length + 1, login: '', firstName: '', lastName: '', email: '', roles: []}
            setUsers([...users, user])
        },
        removeRow() {
            if (users.length <= 1) {
                return;
            }
            setUsers(users.slice(0, -1));
        },
        getData() {
            return users;
        }


    }));

    return (
        <table>
            <thead>
            <tr>
                <th>ID</th>
                <th>Login</th>
                <th>First Name</th>
                <th>Last Name</th>
                <th>Email</th>
                <th>Roles</th>
            </tr>
            </thead>
            <tbody>
            {users.map((user) => <CreateUserTableRow onUpdate={updateUser} key={user.id} id={user.id}/>)}
            </tbody>
        </table>
    );
});

export default CreateUserTable;