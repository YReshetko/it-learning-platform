import React, {useEffect, useState} from 'react';
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
import {get, post} from "../../utils/rest";

const TeacherCourses = () => {
    const [coursesList, setCoursesList] = useState([]);
    const [openForm, setOpenForm] = React.useState(false);
    const [newCourseName, setNewCourseName] = useState('');
    const [newCourseDescription, setNewCourseDescription] = useState('');

    const handleClose = () => {
        setOpenForm(false);
    };
    const handleOpen = () => {
        setOpenForm(true);
    };

    useEffect(() => {
        async function fetch() {
            get("/api/teacher/courses", null)
                .then(response => {
                    setCoursesList(response.courses)
                })
                .catch(error => {
                    console.error('Error:', error)
                })
        }

        fetch();
    }, []);

    const callback = (course) => {
        console.log(course)
    };

    const saveCourse = () => {
        if (newCourseName.trim() === '' || newCourseDescription.trim() === '') {
            console.log("Empty, skip")
            return;
        }

        post("/api/teacher/courses", {
            name: newCourseName,
            description: newCourseDescription,
            seq_no: coursesList.length,
            active: true
        }).then(response => {
            console.log("Course creation:", response)
            setNewCourseName('');
            setNewCourseDescription('');
            setCoursesList([...coursesList, response])
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
                                color="text.primary">Мои курсы</Typography>
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
                                    Добавить курс
                                </Typography>
                                <IconButton aria-label="close" onClick={handleClose} sx={{my: 1, mx: 1}}>
                                    <CloseIcon/>
                                </IconButton>
                            </Toolbar>
                        </div>
                        <div>
                            <TextField fullWidth id="standard-basic" label="Название курса"
                                       onChange={(event) => setNewCourseName(event.currentTarget.value)}
                                       variant="filled">{newCourseName}</TextField>
                        </div>
                        <div>
                            <TextField fullWidth id="outlined-textarea" label="Описание"
                                       onChange={(event) => setNewCourseDescription(event.currentTarget.value)}
                                       variant="filled" multiline rows={10}>{newCourseDescription}</TextField>
                        </div>
                        <div>
                            <Button variant="contained" onClick={saveCourse} sx={{mx: 1, my: 1}}>Сохранить</Button>
                        </div>
                    </Box>
                </Backdrop>
            </div>
            <div>
                <List>
                    {coursesList.map((item) =>
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

export default TeacherCourses;