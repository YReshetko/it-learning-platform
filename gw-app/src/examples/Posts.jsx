import React, {useState} from 'react';
import Post from "./Post";

const Posts = () => {
    const [posts, setPosts] = useState([
        {id: 1, title: "Java", body:"Body"},
        {id: 2, title: "Java", body:"Body"},
        {id: 3, title: "Java", body:"Body"},
    ])
    return (
        <div>
            <h1> Posts list</h1>
            {posts.map(post => <Post post={post} key={post.id}/>)}
        </div>
    );
};

export default Posts;