import React from 'react';
import {Alert, AlertTitle} from "@mui/material";

const UnsupportedLayout = () => {
    return (
        <div>
            <Alert severity="warning">
                <AlertTitle>Внимание</AlertTitle>
                Данная страница ещё не поддерживается! <strong>Разработка - в процессе</strong>
            </Alert>
        </div>
    );
};

export default UnsupportedLayout;