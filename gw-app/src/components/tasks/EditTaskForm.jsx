import React, {useEffect, useState} from 'react';
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import CloseIcon from "@mui/icons-material/Close";
import {Backdrop, Button, Chip, TextField} from "@mui/material";
import Markdown from "../../ui/markdown/Markdown";
import {get, post, remove} from "../../utils/rest";

const EditTaskForm = ({task, open, onClose, onEdit}) => {
    const [tags, setTags] = useState([]);
    const [currentTask, setCurrentTask] = useState(task)

    useEffect(() => {
        async function fetchTags() {
            get("/api/admin/tags", null)
                .then(response => {
                    setTags(response.tags)
                })
                .catch(error => {
                    console.error('Error:', error)
                })
        }
        fetchTags();
    }, [])

    const addTag = (name) => {
        post("/api/admin/tasks/" + task.id + "/tags", {
            name: name,
        }).then(response => {
            if (response.tags) {
                setCurrentTask(response);
                onEdit(response);
            }
        }).catch(error => {
            console.error('Error:', error)
        })
    }
    const removeTag = (name) => {
        remove("/api/admin/tasks/" + task.id + "/tags/" + name).then(response => {
            setCurrentTask(response);
            onEdit(response);
        }).catch(error => {
            console.error('Error:', error)
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
                            Редактировать задачу
                        </Typography>
                        <IconButton aria-label="close" onClick={onClose} sx={{my: 1, mx: 1}}>
                            <CloseIcon/>
                        </IconButton>
                    </Toolbar>
                </div>
                <div>
                    <TextField disabled fullWidth id="standard-basic" label="Название задачи"
                               variant="filled" defaultValue={currentTask.name}/>
                </div>
                <div>
                    <Markdown viewOnly={true} content={currentTask.description}/>
                </div>
                <div>
                    <Button variant="contained" sx={{mx: 1, my: 1}}>Сохранить</Button>
                </div>
                <div>
                    <Typography variant="h6" noWrap component="div" sx={{color: '#000'}}>Тэги задачи</Typography>
                    {currentTask.tags.map((tag) => <Chip sx={{ml: 0.5, mt: 0.5}} label={tag.name} variant="outlined"
                                                       onDelete={() => removeTag(tag.name)}/>)}
                </div>
                <div>
                    <Typography variant="h6" noWrap component="div" sx={{color: '#000'}}>Все тэги</Typography>
                    {tags.map((tag) => <Chip sx={{ml: 0.5, mt: 0.5}} label={tag.name} variant="outlined"
                                             onClick={() => addTag(tag.name)}/>)}
                </div>
            </Box>
        </Backdrop>
    );
};

export default EditTaskForm;