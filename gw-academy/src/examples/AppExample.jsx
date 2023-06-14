import React, {useState} from 'react';
import Counter from "./Counter";
import Posts from "./Posts";
import Button from "./ui/button/Button";

const AppExample = () => {
    const [value, setValue] = useState('hello')
    return (
        <div className="App">
            Working application
            <h1>{value}</h1>
            <input type="text" onChange={event => setValue(event.target.value)}/>
            <Counter count={10}/>
            <Posts/>
            <Button>My button</Button>
        </div>
    );
};

export default AppExample;