import React from 'react';

const CreateUserTableRow = ({id, onUpdate}) => {
    const handleLogin = event => {
        onUpdate(id, 'login', event.target.value)
    }
    const handleFirstName = event => {
        onUpdate(id, 'firstName', event.target.value)
    }
    const handleLastName = event => {
        onUpdate(id, 'lastName', event.target.value)
    }
    const handleEmail = event => {
        onUpdate(id, 'email', event.target.value)
    }
    const handleRoles = event => {
        const roles = event.target.value.split(",").map(str => str.trim())
        onUpdate(id, 'roles', roles)
    }

    return (
        <tr>
            <td>{id}</td>
            <td><input onChange={handleLogin}/></td>
            <td><input onChange={handleFirstName}/></td>
            <td><input onChange={handleLastName}/></td>
            <td><input onChange={handleEmail}/></td>
            <td><input onChange={handleRoles}/></td>
        </tr>
    );
};

export default CreateUserTableRow;