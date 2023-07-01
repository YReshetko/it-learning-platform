import React, {useState} from 'react';
import {Breadcrumbs, Button} from "@mui/material";
import NavigateNextIcon from '@mui/icons-material/NavigateNext';
import TechnologiesList from "./TechnologiesList";
import CategoriesList from "./CategoriesList";
import TopicsList from "./TopicsList";
import Topic from "./Topic";

const TECHNOLOGIES_PAGE = 'technologies';
const CATEGORIES_PAGE = 'categories';
const TOPICS_PAGE = 'topics';
const TOPIC_PAGE = 'topic';

const Technologies = () => {
    const [page, setPage] = useState({
        technology: {
            id: '',
            name: '',
            description: ''
        },
        category: {
            id: '',
            name: '',
            description: ''
        },
        topic: {
            id: '',
            name: '',
            description: ''
        },
        page: TECHNOLOGIES_PAGE
    });

    const backToTechnologiesPage = () => {
        setPage(
            {
                technology: {
                    id: '',
                    name: '',
                    description: ''
                },
                category: {
                    id: '',
                    name: '',
                    description: ''
                },
                topic: {
                    id: '',
                    name: '',
                    description: ''
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
                    id: '',
                    name: '',
                    description: ''
                },
                topic: {
                    id: '',
                    name: '',
                    description: ''
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
                    id: '',
                    name: '',
                    description: ''
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
                    id: '',
                    name: '',
                    description: ''
                },
                topic: {
                    id: '',
                    name: '',
                    description: ''
                },
                page: CATEGORIES_PAGE
            }
        );
    }

    const toCategory = (item) => {
        setPage(
            {
                technology: {
                    id: page.technology.id,
                    name: page.technology.name,
                    description: page.technology.description
                },
                category: {
                    id: item.id,
                    name: item.name,
                    description: item.description
                },
                topic: {
                    id: '',
                    name: '',
                    description: ''
                },
                page: TOPICS_PAGE
            }
        );
    }
    const toTopic = (item) => {
        setPage(
            {
                technology: {
                    id: page.technology.id,
                    name: page.technology.name,
                    description: page.technology.description
                },
                category: {
                    id: page.category.id,
                    name: page.category.name,
                    description: page.category.description
                },
                topic: {
                    id: item.id,
                    name: item.name,
                    description: item.description
                },
                page: TOPIC_PAGE
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
            return (<Topic topic={page.topic}/>);
        } else if (page.page === TOPICS_PAGE) {
            return (<TopicsList category={page.category} callback={toTopic}/>);
        } else if (page.page === CATEGORIES_PAGE) {
            return (<CategoriesList technology={page.technology} callback={toCategory}/>);
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