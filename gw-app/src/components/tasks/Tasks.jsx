import React, {useEffect, useState} from 'react';
import Typography from "@mui/material/Typography";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import AddCircleOutlineOutlinedIcon from "@mui/icons-material/AddCircleOutlineOutlined";
import List from "@mui/material/List";
import ListItemText from "@mui/material/ListItemText";
import EditOutlinedIcon from "@mui/icons-material/EditOutlined";
import {get} from "../../utils/rest";
import NewTaskForm from "./NewTaskForm";
import EditTaskForm from "./EditTaskForm";
import ListItem from "@mui/material/ListItem";

const Tasks = () => {
    const [tasksList, setTasksList] = useState([]);
    const [taskToEdit, setTaskToEdit] = useState(null);
    const [openNewTaskForm, setOpenNewTaskForm] = React.useState(false);
    const handleNewTaskFormClose = () => {
        setOpenNewTaskForm(false);
    };
    const handleNewTaskFormOpen = () => {
        setOpenNewTaskForm(true);
    };


    useEffect(() => {
        async function fetch() {
            get("/api/admin/tasks", null)
                .then(response => {
                    setTasksList(response.tasks);
                })
                .catch(error => {
                    console.error('Error:', error);
                })
        }

        fetch();
    }, []);

    const addNewTask = (task) => {
        setTasksList([...tasksList, task])
    }

    const editTask = (task) => {
        setTasksList(tasksList.map((tsk) => task.id === tsk.id ? task : tsk))
    }

    const getEditForm = () => {
        if (taskToEdit === null) {
            return (<div/>);
        } else {
            return (<EditTaskForm open={true} onClose={() => setTaskToEdit(null)} task={taskToEdit} onEdit={editTask}/>);
        }
    }

    return (
        <div>
            <div>
                <Toolbar>
                    <Typography variant="h4" noWrap component="div"
                                sx={{color: '#000', mx: 1, my: 1, flexGrow: 1}}
                                color="text.primary">Задачи</Typography>
                    <IconButton aria-label="close" onClick={handleNewTaskFormOpen} sx={{my: 1, mx: 1}}>
                        <AddCircleOutlineOutlinedIcon/>
                    </IconButton>
                </Toolbar>
            </div>
            <div>
                <NewTaskForm open={openNewTaskForm} onClose={handleNewTaskFormClose} onAdd={addNewTask}/>
                {getEditForm()}
            </div>
            <div>
                <List>
                    {tasksList.map((item) =>
                        <ListItem
                            secondaryAction={
                                <IconButton onClick={() => setTaskToEdit(item)} edge="end" aria-label="edit">
                                    <EditOutlinedIcon/>
                                </IconButton>
                            }>
                            <ListItemText primary={item.name} sx={{flexGrow: 1}}/>
                        </ListItem>)}
                </List>
            </div>
        </div>
    );
};

export default Tasks;