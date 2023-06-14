import React, {useState} from 'react';

const Counter = (props) => {
    const [count, setCount] = useState(props.count);
    function increment(){
        setCount(count+1)
    }
    function decrement(){
        setCount(count-1)
    }
    return (
        <div>
            <h1>{count}</h1>
            <button onClick={increment}>Increnment</button>
            <button onClick={decrement}>Decrement</button>
        </div>
    );
};

export default Counter;