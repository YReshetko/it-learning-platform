import React, {useEffect, useState} from "react";
import './styles/post.css'
import UserForm from "./components/users/UserForm";
import * as auth from './utils/auth.js';
/*import Header from "./components/header/Header";*/
import {get} from "./utils/rest";
import Workspace from "./components/workspace/Workspace";

function App() {
    auth.verifyToken();

    const [me, setMe] = useState({
        id: '',
        firstName: 'John',
        lastName: 'Doe',
        roles: ["ADMIN", "STUDENT", "MANAGER", "TEACHER"]
    })
    useEffect(() => {
            async function fetch() {
                get("/api/self", null)
                    .then(response => {
                        setMe({
                            id: response.id,
                            firstName: response.first_name,
                            lastName: response.last_name,
                            roles: response.roles
                        })
                    })
                    .catch(error => {
                        console.error('Error:', error)
                    })
            }
            fetch();
        },
        []);

    return (
        <div className="App">
            {/*<Header lastName={me.lastName} firstName={me.firstName}></Header>*/}
            {/*<UserForm/>*/}
            <Workspace user={me}/>
        </div>
    );
}

export default App;
