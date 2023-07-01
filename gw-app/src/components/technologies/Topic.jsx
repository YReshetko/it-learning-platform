import React from 'react';
import Typography from "@mui/material/Typography";

const Topic = ({topic}) => {
    return (
        <div>
            <Typography variant="h3" noWrap component="div" sx={{color: '#000'}}>{topic.name}</Typography>
            <Typography variant="body1" noWrap component="div" sx={{color: '#000'}}>
                    <pre style={{fontFamily: 'inherit'}}>
                        {topic.description}
                    </pre>
            </Typography>
        </div>
    );
};

export default Topic;