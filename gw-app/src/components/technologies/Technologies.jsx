import React, {useState} from 'react';
import {Breadcrumbs, Button} from "@mui/material";
import NavigateNextIcon from '@mui/icons-material/NavigateNext';
import TechnologiesList from "./TechnologiesList";
import CategoriesList from "./CategoriesList";

const TECHNOLOGIES_PAGE = 'technologies';
const CATEGORIES_PAGE = 'categories';
const TOPICS_PAGE = 'topics';
const TOPIC_PAGE = 'topic';

const Technologies = () => {
    const [page, setPage] = useState({
        technology: {
            name: '',
            id: ''
        },
        category: {
            name: '',
            id: ''
        },
        topic: {
            name: '',
            id: ''
        },
        page: TECHNOLOGIES_PAGE
    });

    const backToTechnologiesPage = () => {
        setPage(
            {
                technology: {
                    name: '',
                    id: ''
                },
                category: {
                    name: '',
                    id: ''
                },
                topic: {
                    name: '',
                    id: ''
                },
                page: TECHNOLOGIES_PAGE
            }
        );
    }

    const backToCategoriesPage = () => {
        setPage(
            {
                ...page,
                category: {
                    name: '',
                    id: ''
                },
                topic: {
                    name: '',
                    id: ''
                },
                page: CATEGORIES_PAGE
            }
        );
    }

    const backToTopicsPage = () => {
        setPage(
            {
                ...page,
                topic: {
                    name: '',
                    id: ''
                },
                page: TOPICS_PAGE
            }
        );
    }

    const toTechnology = (item) => {
        setPage(
            {
                technology: {
                    id: item.id,
                    name: item.name,
                    description: item.description
                },
                category: {
                    name: '',
                    id: ''
                },
                topic: {
                    name: '',
                    id: ''
                },
                page: CATEGORIES_PAGE
            }
        );
    }

    const getTopicBreadcrumbs = () => {
        return (
            <Breadcrumbs separator={<NavigateNextIcon fontSize="small"/>} aria-label="breadcrumb">
                <Button variant="text" onClick={backToTechnologiesPage}>{page.technology.name}</Button>
                <Button variant="text" onClick={backToCategoriesPage}>{page.category.name}</Button>
                <Button variant="text" onClick={backToTopicsPage}>{page.topic.name}</Button>
            </Breadcrumbs>
        )
    }
    const getTopicsBreadcrumbs = () => {
        return (
            <Breadcrumbs separator={<NavigateNextIcon fontSize="small"/>} aria-label="breadcrumb">
                <Button variant="text" onClick={backToTechnologiesPage}>{page.technology.name}</Button>
                <Button variant="text" onClick={backToCategoriesPage}>{page.category.name}</Button>
            </Breadcrumbs>
        )
    }
    const getCategoriesBreadcrumbs = () => {
        return (
            <Breadcrumbs separator={<NavigateNextIcon fontSize="small"/>} aria-label="breadcrumb">
                <Button variant="text" onClick={backToTechnologiesPage}>{page.technology.name}</Button>
            </Breadcrumbs>
        )
    }

    const getBreadcrumbs = () => {
        if (page.page === TOPIC_PAGE) {
            return getTopicBreadcrumbs();
        } else if (page.page === TOPICS_PAGE) {
            return getTopicsBreadcrumbs();
        } else if (page.page === CATEGORIES_PAGE) {
            return getCategoriesBreadcrumbs();
        } else {
            return (<Breadcrumbs separator={<NavigateNextIcon fontSize="small"/>} aria-label="breadcrumb"/>);
        }
    }

    const getPageList = () => {
        if (page.page === TOPIC_PAGE) {
            return (<TechnologiesList/>);
        } else if (page.page === TOPICS_PAGE) {
            return (<CategoriesList/>);
        } else if (page.page === CATEGORIES_PAGE) {
            return (<CategoriesList/>);
        } else {
            return (<TechnologiesList callback={toTechnology}/>);
        }
    }

    return (
        <div>
            <div>
                {getBreadcrumbs()}
            </div>
            {getPageList()}
        </div>
    );
};

export default Technologies;