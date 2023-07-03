import React, {useEffect, useState} from 'react';
import ReactMarkdown from "react-markdown";
import {Container, TextField, ToggleButton, ToggleButtonGroup} from "@mui/material";
import remarkGfm from 'remark-gfm';
import {Prism as SyntaxHighlighter} from "react-syntax-highlighter";

const Markdown = ({content, viewOnly, onChange}) => {
    const [currentContent, setCurrentContent] = useState(content);
    const [view, setView] = useState('edit')

    useEffect(() => {
        if (viewOnly) {
            setView('preview')
        }
    }, []);

    const changeContent = (event) => {
        setCurrentContent(event.currentTarget.value)
        onChange(event.currentTarget.value)
    };

    const getView = () => {
        if (view === 'edit') {
            return (
                <TextField fullWidth id="outlined-textarea"
                           onChange={changeContent}
                           variant="outlined" multiline rows={15} defaultValue={currentContent}/>
            );
        } else {
            return (
                <Container>
                    <ReactMarkdown
                        remarkPlugins={[remarkGfm]}
                        children={currentContent}
                        components={{
                            code({node, inline, className, children, ...props}) {
                                const match = /language-(\w+)/.exec(className || "");
                                return !inline && match ? (
                                    <SyntaxHighlighter
                                        children={String(children).replace(/\n$/, "")}
                                        language={match[1]}
                                        {...props}
                                    />
                                ) : (
                                    <code className={className} {...props}>
                                        {children}
                                    </code>
                                );
                            },
                        }}
                    />
                </Container>
            );
        }
    }
    const handleChange = (event, newView) => {
        setView(newView);
    };

    const displayControle = () => {
        if (viewOnly) {
            return (<div/>);
        } else {
            return (
                <ToggleButtonGroup
                    color="primary"
                    value={view}
                    exclusive
                    onChange={handleChange}
                    aria-label="control"
                >
                    <ToggleButton value="edit">Редактировать</ToggleButton>
                    <ToggleButton value="preview">Предпросмотр</ToggleButton>
                </ToggleButtonGroup>
            );
        }
    }
    return (
        <div>
            <Container maxWidth="md" sx={{backgroundColor: '#FAFAFA', color: '#000'}}>
                <div>
                    {getView()}
                </div>
            </Container>
            {displayControle()}
        </div>
    );
};

export default Markdown;