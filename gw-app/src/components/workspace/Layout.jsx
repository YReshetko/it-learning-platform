import React from 'react';
import UserForm from "../users/UserForm";
import UnsupportedLayout from "../unsupported/UnsupportedLayout";
import Technologies from "../technologies/Technologies";
import Tags from "../tags/Tags";
import Tasks from "../tasks/Tasks";
import TeacherCourses from "../courses/TeacherCourses";

const layoutsMapping = {
    "admin-add-user": (<UserForm/>),
    "admin-technologies": (<Technologies/>),
    "admin-tags": (<Tags/>),
    "admin-tasks": (<Tasks/>),

    "teacher-my-courses": (<TeacherCourses/>)
    /*"admin-all-users": (<UserList/>)*/
}

const getLayout = (layout) => {
    let out = layoutsMapping[layout];
    if (out !== undefined) {
        return out;
    } else {
        return (<UnsupportedLayout/>);
    }
}

const Layout = ({layout}) => {
    return (
        <div>
            {getLayout(layout)}
        </div>
    );
};

export default Layout;