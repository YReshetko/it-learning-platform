import React, {useEffect, useState} from 'react';
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import {Backdrop, Button, TextField} from "@mui/material";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";
import List from "@mui/material/List";
import {get, post} from "../../utils/rest";
import CloseIcon from '@mui/icons-material/Close';
import EditOutlinedIcon from '@mui/icons-material/EditOutlined';
import DeleteOutlineOutlinedIcon from '@mui/icons-material/DeleteOutlineOutlined';
import AddCircleOutlineOutlinedIcon from '@mui/icons-material/AddCircleOutlineOutlined';
import IconButton from "@mui/material/IconButton";
import Toolbar from "@mui/material/Toolbar";

const TechnologiesList = ({callback}) => {
    const [technologiesList, setTechnologiesList] = useState([]);
    const [openForm, setOpenForm] = React.useState(false);
    const [newTechnologyName, setNewTechnologyName] = useState('');
    const [newTechnologyDescription, setNewTechnologyDescription] = useState('');

    const handleClose = () => {
        setOpenForm(false);
    };
    const handleOpen = () => {
        setOpenForm(true);
    };
    useEffect(() => {
        async function fetch() {
            get("/api/admin/technologies", null)
                .then(response => {
                    setTechnologiesList(response.technologies)
                })
                .catch(error => {
                    console.error('Error:', error)
                })
        }

        fetch();
    }, []);

    const saveTechnology = () => {
        if (newTechnologyName.trim() === '' || newTechnologyDescription.trim() === '') {
            console.log("Empty, skip")
            return;
        }

        post("/api/admin/technology", {
            name: newTechnologyName,
            description: newTechnologyDescription
        }).then(response => {
            console.log("Technology creation status:", response)
            setNewTechnologyName('');
            setNewTechnologyDescription('');
        }).catch(error => {
            console.error('Error:', error)
        })
    }

    return (
        <div>
            <div>
                <Toolbar>
                    <Typography variant="h4" noWrap component="div"
                                sx={{color: '#000', mx: 1, my: 1, flexGrow: 1}}
                                color="text.primary">Технологии</Typography>
                    <IconButton aria-label="close" onClick={handleOpen} sx={{my: 1, mx: 1}}>
                        <AddCircleOutlineOutlinedIcon/>
                    </IconButton>
                </Toolbar>
            </div>
            <div>
                <Backdrop sx={{color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1}} open={openForm}>
                    <Box component="form"
                         sx={{'& .MuiTextField-root': {m: 1, width: '100ch'}, backgroundColor: '#EEE'}} noValidate
                         autoComplete="off">
                        <div>
                            <Toolbar>
                                <Typography variant="h6" noWrap component="div"
                                            sx={{color: '#000', mx: 1, my: 1, flexGrow: 1}}>
                                    Добавить технологию
                                </Typography>
                                <IconButton aria-label="close" onClick={handleClose} sx={{my: 1, mx: 1}}>
                                    <CloseIcon/>
                                </IconButton>
                            </Toolbar>
                        </div>
                        <div>
                            <TextField fullWidth id="standard-basic" label="Название технологии"
                                       onChange={(event) => setNewTechnologyName(event.currentTarget.value)}
                                       variant="filled">{newTechnologyName}</TextField>
                        </div>
                        <div>
                            <TextField fullWidth id="outlined-textarea" label="Описание"
                                       onChange={(event) => setNewTechnologyDescription(event.currentTarget.value)}
                                       variant="filled" multiline rows={10}>{newTechnologyDescription}</TextField>
                        </div>
                        <div>
                            <Button variant="contained" onClick={saveTechnology} sx={{mx: 1, my: 1}}>Сохранить</Button>
                        </div>
                    </Box>
                </Backdrop>
            </div>
            <div>
                <List>
                    {technologiesList.map((item) =>
                        <ListItemButton onClick={(e) => callback(item)} key={item.id}>
                            <ListItemText primary={item.name} sx={{flexGrow: 1}}/>
                            <IconButton disabled aria-label="edit">
                                <EditOutlinedIcon/>
                            </IconButton>
                            <IconButton disabled aria-label="delete">
                                <DeleteOutlineOutlinedIcon/>
                            </IconButton>
                        </ListItemButton>)}
                </List>
            </div>
        </div>
    );
};

export default TechnologiesList;