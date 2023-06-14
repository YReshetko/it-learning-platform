import React from 'react';

const Post = (props) => {
    return (
        <div className="post">
            <div className="post_content">
                <strong>{props.post.id}. {props.post.title}</strong>
                <div>
                    {props.post.body}
                </div>
                <div className="post_btns">
                    <button>DELETE</button>
                </div>
            </div>
        </div>
    );
};

export default Post;