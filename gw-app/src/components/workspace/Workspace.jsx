import React, {useState} from 'react';
import Toolbar from "@mui/material/Toolbar";
import ListItemText from "@mui/material/ListItemText";
import ListItemButton from "@mui/material/ListItemButton";
import Divider from "@mui/material/Divider";
import List from "@mui/material/List";
import Box from "@mui/material/Box";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import Typography from "@mui/material/Typography";
import AppBar from "@mui/material/AppBar";
import Drawer from "@mui/material/Drawer";
import Layout from "./Layout";
import {Button} from "@mui/material";
import * as rest from "../../utils/rest";
import * as auth from "../../utils/auth";

const menuItems = {
    "ADMIN": {
        label: "Администратор",
        menu: [
            {name: "Добавить пользователей", label: "admin-add-user"},
            {name: "Все пользователи", label: "admin-all-users"},
            {name: "Технологии", label: "admin-technologies"},
        ]
    },
    "TEACHER": {
        label: "Преподаватель",
        menu: [
            {name: "Все курсы", label: "teacher-all-courses"},
            {name: "Мои курсы", label: "teacher-my-courses"},
            {name: "Мои группы", label: "teacher-my-groups"},
            {name: "Создать группу", label: "teacher-ceate-group"}
        ]
    },
    "STUDENT": {
        label: "Студент",
        menu: [
            {name: "Мои курсы", label: "student-my-courses"}
        ]
    },
    "MANAGER": {
        label: "Методист",
        menu: [
            {name: "Мои группы", label: "manager-my-groups"}
        ]
    }
}
const drawerWidth = 240;
const appBarWidths = [240, 0];
const Workspace = ({user}) => {
    const [menuOpen, setMenuOpen] = useState(true);
    const [appBarWidth, setAppBarWidth] = useState(appBarWidths[0]);
    const [layout, setLayout] = useState('admin-add-user');

    const handleMenuOpen = () => {
        setMenuOpen(!menuOpen);
        if (appBarWidth === appBarWidths[0]) {
            setAppBarWidth(appBarWidths[1])
        } else {
            setAppBarWidth(appBarWidths[0])
        }
    };

    const openLayout = (layout) => {
        setLayout(layout)
        console.log("Opening layout", layout)
    }

    const logout = () => {
        rest.post("/api/logout")
            .then(response => {
                auth.clean();
                window.location.replace("/");
            })
            .catch(error => {
                console.error('Error:', error)
            });
    }

    return (
        <div>
            <Box sx={{display: 'flex'}}>
                {/*<AppBar position="fixed" sx={{ width: { sm: `calc(100% - ${drawerWidth}px)` }, ml: { sm: `${drawerWidth}px` }}}>*/}
                <AppBar position="fixed" sx={{width: {sm: `calc(100% - ${appBarWidth}px)`} , ml: { sm: `${appBarWidth}px` }}}>
                    <Toolbar>
                        <IconButton color="inherit" onClick={handleMenuOpen}>
                            <MenuIcon/>
                        </IconButton>
                        <Typography variant="h6" noWrap component="div" sx={{ flexGrow: 1 }}>
                            {user.firstName} {user.lastName}
                        </Typography>
                        <Button onClick={logout} color="inherit">Выход</Button>
                    </Toolbar>
                </AppBar>
                <Box component="nav" sx={{width: {sm: drawerWidth}, flexShrink: {sm: 0}}}>
                    {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
                    <Drawer
                        variant="persistent"
                        sx={{
                            display: {xs: 'none', sm: 'block'},
                            '& .MuiDrawer-paper': {boxSizing: 'border-box', width: drawerWidth},
                        }}
                        ModalProps={{
                            keepMounted: false, // Better open performance on mobile.
                        }}
                        open={menuOpen}
                    >
                        {createMenuItems(user.roles, openLayout)}
                    </Drawer>
                </Box>
                <Box
                    component="main"
                    sx={{ flexGrow: 1, p: 3, width: { sm: `calc(100% - ${appBarWidth}px)` } }}
                >
                    <Toolbar />
                    <Layout layout={layout}/>
                </Box>
            </Box>
        </div>
    );
};

function createMenuItems(userRoles, callback) {
    let userMenuItems = userRoles.map(userRole => menuItems[userRole]);

    return (
        <div>
            <Toolbar sx={{backgroundColor: '#1976d2'}}>
                <Typography sx={{color: 'white'}} variant="h6" noWrap component="div">
                    Меню
                </Typography>
            </Toolbar>

            {userMenuItems.map((item) => createMenuBlock(item, callback))}
        </div>
    )
}

function createMenuBlock(block, callback) {
    return (
        <div key={block.label}>
            <Divider>{block.label}</Divider>
            <List>
                {block.menu.map((menuItem) =>
                    <ListItemButton onClick={(e) => callback(menuItem.label)} key={menuItem.label}>
                        <ListItemText primary={menuItem.name}/>
                    </ListItemButton>)}
            </List>
        </div>
    )
}

export default Workspace;