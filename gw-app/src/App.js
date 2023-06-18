import React, {useEffect, useState} from "react";
import './styles/post.css'
import UserForm from "./components/users/UserForm";
import * as auth from './utils/auth.js';
import Header from "./components/header/Header";
import {get} from "./utils/rest";

function App() {
    auth.verifyToken();

    const [me, setMe] = useState({id: '', firstName: '', lastName: ''})
    useEffect(() => get("/api/self", null)
            .then(response => {
                setMe({
                    id: response.id,
                    firstName: response.first_name,
                    lastName: response.last_name
                })
            })
            .catch(error => {
                console.error('Error:', error)
            }),
        []);

    return (
        <div className="App">
            <Header lastName={me.lastName} firstName={me.firstName}></Header>
            <UserForm/>
        </div>
    );
}

export default App;
