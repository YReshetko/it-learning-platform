import React, {useEffect, useState} from 'react';
import {get, post} from "../../utils/rest";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import AddCircleOutlineOutlinedIcon from "@mui/icons-material/AddCircleOutlineOutlined";
import {Backdrop, Button, TextField} from "@mui/material";
import Box from "@mui/material/Box";
import CloseIcon from "@mui/icons-material/Close";
import List from "@mui/material/List";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";
import EditOutlinedIcon from "@mui/icons-material/EditOutlined";
import DeleteOutlineOutlinedIcon from "@mui/icons-material/DeleteOutlineOutlined";

const CategoriesList = ({technology, callback}) => {
    const [categoriesList, setCategoriesList] = useState([]);
    const [openForm, setOpenForm] = React.useState(false);
    const [newCategoryName, setNewCategoryName] = useState('');
    const [newCategoryDescription, setNewCategoryDescription] = useState('');

    const handleClose = () => {
        setOpenForm(false);
    };
    const handleOpen = () => {
        setOpenForm(true);
    };

    useEffect(() => {
        async function fetch() {
            get("/api/admin/technologies/" + technology.id + "/categories", null)
                .then(response => {
                    setCategoriesList(response.categories)
                })
                .catch(error => {
                    console.error('Error:', error)
                })
        }

        fetch();
    }, []);

    const saveCategory = () => {
        if (newCategoryName.trim() === '' || newCategoryDescription.trim() === '') {
            console.log("Empty, skip")
            return;
        }

        post("/api/admin/technologies/" + technology.id + "/category", {
            name: newCategoryName,
            description: newCategoryDescription
        }).then(response => {
            setNewCategoryName('');
            setNewCategoryDescription('');
            setCategoriesList([...categoriesList, response])
        }).catch(error => {
            console.error('Error:', error)
        })
    }

    return (
        <div>
            <div>
                <Typography variant="h3" noWrap component="div" sx={{color: '#000'}}>{technology.name}</Typography>
                <Typography variant="body1" noWrap component="div" sx={{color: '#000'}}>
                    <pre style={{fontFamily: 'inherit'}}>
                        {technology.description}
                    </pre>
                </Typography>
            </div>
            <div>
                <Toolbar>
                    <Typography variant="h4" noWrap component="div"
                                sx={{color: '#000', mx: 1, my: 1, flexGrow: 1}}
                                color="text.primary">Категории</Typography>
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
                                    Добавить категорию
                                </Typography>
                                <IconButton aria-label="close" onClick={handleClose} sx={{my: 1, mx: 1}}>
                                    <CloseIcon/>
                                </IconButton>
                            </Toolbar>
                        </div>
                        <div>
                            <TextField fullWidth id="standard-basic" label="Название категории"
                                       onChange={(event) => setNewCategoryName(event.currentTarget.value)}
                                       variant="filled">{newCategoryName}</TextField>
                        </div>
                        <div>
                            <TextField fullWidth id="outlined-textarea" label="Описание"
                                       onChange={(event) => setNewCategoryDescription(event.currentTarget.value)}
                                       variant="filled" multiline rows={10}>{newCategoryDescription}</TextField>
                        </div>
                        <div>
                            <Button variant="contained" onClick={saveCategory} sx={{mx: 1, my: 1}}>Сохранить</Button>
                        </div>
                    </Box>
                </Backdrop>
            </div>
            <div>
                <List>
                    {categoriesList.map((item) =>
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

export default CategoriesList;