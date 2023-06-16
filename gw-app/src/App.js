import React from "react";
import './styles/post.css'
import UserForm from "./components/users/UserForm";
import * as auth from './utils/auth.js';

function App() {
    auth.verifyToken();
    return (
        <div className="App">
            <UserForm/>
        </div>
    );
}

export default App;
