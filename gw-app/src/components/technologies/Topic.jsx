import React, {useEffect, useState} from 'react';
import Typography from "@mui/material/Typography";
import {get, post, remove} from "../../utils/rest";
import {Chip} from "@mui/material";

const Topic = ({topic}) => {
    const [tags, setTags] = useState([]);
    const [fullTopic, setFullTopic] = useState({tags:[]});

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

        async function fetchTopic() {
            get("/api/admin/topics/" + topic.id, null)
                .then(response => {
                    setFullTopic(response)
                })
                .catch(error => {
                    console.error('Error:', error)
                })
        }

        fetchTags();
        fetchTopic();
    }, [])

    const addTag = (name) => {
        post("/api/admin/topics/" + fullTopic.id + "/tags", {
            name: name,
        }).then(response => {
            if (response.tags) {
                setFullTopic(response);
            }
        }).catch(error => {
            console.error('Error:', error)
        })
    }
    const removeTag = (name) => {
        remove("/api/admin/topics/" + fullTopic.id + "/tags/" + name).then(response => {
            setFullTopic(response);
        }).catch(error => {
            console.error('Error:', error)
        })
    }

    return (
        <div>
            <div>
                <Typography variant="h3" noWrap component="div" sx={{color: '#000'}}>{topic.name}</Typography>
                <Typography variant="body1" noWrap component="div" sx={{color: '#000'}}>
                    <pre style={{fontFamily: 'inherit'}}>
                        {topic.description}
                    </pre>
                </Typography>
            </div>
            <div>
                <Typography variant="h6" noWrap component="div" sx={{color: '#000'}}>Тэги темы</Typography>
                {fullTopic.tags.map((tag) => <Chip sx={{ml: 0.5, mt: 0.5}} label={tag.name} variant="outlined"
                                         onDelete={() => removeTag(tag.name)}/>)}
            </div>
            <div>
                <Typography variant="h6" noWrap component="div" sx={{color: '#000'}}>Все тэги</Typography>
                {tags.map((tag) => <Chip sx={{ml: 0.5, mt: 0.5}} label={tag.name} variant="outlined"
                                         onClick={() => addTag(tag.name)}/>)}
            </div>
        </div>
    );
};

export default Topic;