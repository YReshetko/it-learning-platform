import React, {useEffect, useState} from 'react';
import {get} from "../../utils/rest";
import List from "@mui/material/List";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";
import Typography from "@mui/material/Typography";

const Course = ({courseId, onClick}) => {
    const [technologies, setTechnologies] = useState([])
    const [course, setCourse] = useState({})
    useEffect(() => {
        async function fetchTechnologies() {
            get("/api/teacher/technologies", null)
                .then(response => {
                    setTechnologies(response.technologies)
                })
                .catch(error => {
                    console.error('Error:', error)
                })
        }

        async function fetchCourse() {
            get("/api/teacher/courses/" + courseId, null)
                .then(response => {
                    setCourse(response)
                })
                .catch(error => {
                    console.error('Error:', error)
                })
        }

        fetchTechnologies();
        fetchCourse();
    }, []);

    const renderTopics = () => {
        if (course.topics === undefined) {
            return (<div></div>);
        } else {
            <List>
                {course.topics.map((item) =>
                    <ListItemButton>
                        <ListItemText primary={item.name} sx={{flexGrow: 1}}/>
                    </ListItemButton>)}
            </List>
        }
    }
    return (
        <div>
            <div>
                <Typography variant="h4" noWrap component="div"
                            sx={{color: '#000', mx: 1, my: 1, flexGrow: 1}}
                            color="text.primary">{course.name}</Typography>
            </div>
            <div>
                {renderTopics()}
            </div>

            <div>
                <Typography variant="h4" noWrap component="div"
                            sx={{color: '#000', mx: 1, my: 1, flexGrow: 1}}
                            color="text.primary">Технологии и темы</Typography>
            </div>
            <div>
                <List>
                    {technologies.map((item) =>
                        <ListItemButton>
                            <ListItemText primary={item.technology.name} sx={{flexGrow: 1}}/>
                        </ListItemButton>)}
                </List>
            </div>
        </div>
    );
};

export default Course;