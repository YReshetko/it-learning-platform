import classes from './Button.module.css'

const Button = ({ onClick, children }) => {
    const handleClick = (event) => {
        onClick(event.target.value);
    }
    return (
        <button className={classes.btn} onClick={handleClick}>
            {children}
        </button>
    );
};

export default Button;