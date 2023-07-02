import React, {useEffect, useState} from 'react';
import Typography from "@mui/material/Typography";
import {Button, Chip, TextField} from "@mui/material";
import {get, post, remove} from "../../utils/rest";

const Tags = () => {
    const [tags, setTags] = useState([]);
    const [newTag, setNewTag] = useState('');

    useEffect(() => {
        async function fetch() {
            get("/api/admin/tags", null)
                .then(response => {
                    setTags(response.tags)
                })
                .catch(error => {
                    console.error('Error:', error)
                })
        }

        fetch();
    }, [])

    const handleCreate = () => {
        post("/api/admin/tags", {
            name: newTag,
        }).then(response => {
            setNewTag('');
            setTags([...tags, {name: response.name}])
        }).catch(error => {
            console.error('Error:', error)
        })
    }

    const handleDelete = (name) => {
        remove("/api/admin/tags/" + name, {
            name: newTag,
        }).then(response => {
            setNewTag('');
            setTags(tags.filter((tag) => tag.name !== name))
        }).catch(error => {
            console.error('Error:', error)
        })
    }

    return (
        <div>
            <div>
                <Typography variant="h3" noWrap component="div" sx={{color: '#000'}}>Тэги</Typography>
            </div>
            <div>
                <TextField fullWidth noWrap sx={{flexGrow: 1}} id="standard-basic" label="Имя тэга"
                           onChange={(event) => setNewTag(event.currentTarget.value)}
                           variant="filled">{newTag}</TextField>
                <Button variant="contained" onClick={handleCreate} sx={{mx: 1, my: 1}}>Создать</Button>
            </div>
            <div>
                {tags.map((tag) => <Chip sx={{ml: 0.5, mt: 0.5}} label={tag.name} variant="outlined"
                                         onDelete={() => handleDelete(tag.name)}/>)}
            </div>
        </div>
    );
};

export default Tags;