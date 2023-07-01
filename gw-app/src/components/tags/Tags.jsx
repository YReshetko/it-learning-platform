import React, {useState} from 'react';
import Typography from "@mui/material/Typography";
import {Button, Chip, TextField} from "@mui/material";

const Tags = () => {
    const [tags, setTags] = useState([
        {name: 'java'},
        {name: 'примитивы'},
        {name: 'if'},
        {name: 'switch'},
        {name: 'циклы'},
        {name: 'for'},
        {name: 'do-while'},
        {name: 'while'}
    ]);

    const [newTag, setNewTag] = useState('');

    const handleDelete = (name) => {
        console.log("delete tag: ", name);
    }

    const handleCreate = () => {
        setTags([...tags, {name: newTag}])
    }

    return (
        <div>
            <div>
                <Typography variant="h3" noWrap component="div" sx={{color: '#000'}}>Тэги</Typography>
            </div>
            <div>
                <TextField fullWidth noWrap sx={{ flexGrow: 1 }} id="standard-basic" label="Имя тэга"
                           onChange={(event) => setNewTag(event.currentTarget.value)}
                           variant="filled">{newTag}</TextField>
                <Button variant="contained" onClick={handleCreate} sx={{mx: 1, my: 1}}>Создать</Button>
            </div>
            <div>
                {tags.map((tag) => <Chip sx={{ml: 0.5, mt: 0.5}} label={tag.name} variant="outlined" onDelete={() => handleDelete(tag.name)}/>)}
            </div>
        </div>
    );
};

export default Tags;