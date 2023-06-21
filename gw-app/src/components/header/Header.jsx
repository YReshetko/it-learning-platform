import React from 'react';
import * as auth from "../../utils/auth.js"
import {Button, Stack} from "@mui/material";
import DehazeIcon from '@mui/icons-material/Dehaze';

const Header = ({firstName, lastName}) => {
    const logout = (event) => {
        auth.clean();
    }
    const openMenu = (even) => {
        console.log("Open menu")
    }

    let userName = lastName + ' ' + firstName
    return (
        <div>
            <Stack direction={{xs: 'column', sm: 'row'}} spacing={{xs: 1, sm: 2, md: 4}}>
                <Button onClick={openMenu} variant="outlined"><DehazeIcon/></Button>
                <Button onClick={logout} variant="outlined">Выйти</Button>
                <h2 style={{display: "inline-block"}}>{userName}</h2>
            </Stack>
        </div>
    );
};

export default Header;