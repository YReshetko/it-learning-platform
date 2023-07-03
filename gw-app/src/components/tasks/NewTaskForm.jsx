import React, {useState} from 'react';
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import CloseIcon from "@mui/icons-material/Close";
import {Backdrop, Button, TextField} from "@mui/material";
import {post} from "../../utils/rest";
import Markdown from "../../ui/markdown/Markdown";

const NewTaskForm = ({open, onClose, onAdd}) => {
    const [taskName, setTaskName] = useState('');
    const [taskDescription, setTaskDescription] = useState('');


    const saveTask = () => {
        if (taskName.trim() === '' || taskDescription.trim() === '') {
            console.log("Empty, skip")
            return;
        }

        post("/api/admin/tasks", {
            name: taskName,
            description: taskDescription,
            active: true,
            seq_no: 1

        }).then(response => {
            setTaskName('');
            setTaskDescription('');
            onAdd(response);
        }).catch(error => {
            console.error('Error:', error);
        })
    }
    return (
        <Backdrop sx={{color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1}} open={open}>
            <Box component="form"
                 sx={{'& .MuiTextField-root': {m: 1, width: '75ch'}, backgroundColor: '#EEE'}} noValidate
                 autoComplete="off">
                <div>
                    <Toolbar>
                        <Typography variant="h6" noWrap component="div"
                                    sx={{color: '#000', mx: 1, my: 1, flexGrow: 1}}>
                            Добавить задачу
                        </Typography>
                        <IconButton aria-label="close" onClick={onClose} sx={{my: 1, mx: 1}}>
                            <CloseIcon/>
                        </IconButton>
                    </Toolbar>
                </div>
                <div>
                    <TextField fullWidth id="standard-basic" label="Название задачи"
                               onChange={(event) => setTaskName(event.currentTarget.value)}
                               variant="filled">{taskName}</TextField>
                </div>
                <div>
                    <Markdown content={taskDescription} onChange={(data) => setTaskDescription(data)}/>
                </div>
                <div>
                    <Button variant="contained" onClick={saveTask} sx={{mx: 1, my: 1}}>Сохранить</Button>
                </div>
            </Box>
        </Backdrop>
    );
};

export default NewTaskForm;