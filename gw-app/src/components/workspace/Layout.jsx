import React from 'react';
import UserForm from "../users/UserForm";
import UserList from "../users/UserList";
import UnsupportedLayout from "../unsupported/UnsupportedLayout";

const layoutsMapping = {
    "admin-add-user": (<UserForm/>),
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