import React from 'react';
import classes from "./Label.module.css";

const Label = ({children}) => {
    return (<div className={classes.lbl}>{children}</div>);
};

export default Label;