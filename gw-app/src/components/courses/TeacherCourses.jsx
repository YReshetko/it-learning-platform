import React, {useState} from 'react';
import CoursesList from "./CoursesList";
import EditOutlinedIcon from "@mui/icons-material/EditOutlined";
import ArrowCircleLeftOutlinedIcon from '@mui/icons-material/ArrowCircleLeftOutlined';
import IconButton from "@mui/material/IconButton";
import Course from "./Course";

const TeacherCourses = () => {
    const [stage, setStage] = useState('courses'); // course // topic
    const [courseId, setCourseId] = useState('');
    const [topicId, setTopicId] = useState('');

    const goToCourse = (courseId) => {
        setCourseId(courseId)
        setStage('course')
    }

    const goToTopic = (topicId) => {
        setTopicId(topicId)
        setStage('topic')
    }

    const backToCourse = () => {
        setTopicId('')
        setStage('course')
    }

    const backToCourses = () => {
        setCourseId('')
        setStage('courses')
    }

    const resolveStage = () => {
        switch (stage) {
            case 'courses':
                return (
                    <CoursesList onClick={goToCourse}/>
                );
            case 'course':
                return (
                    <div>
                        {/*back button*/}
                        <IconButton onClick={backToCourses} aria-label="edit">
                            <ArrowCircleLeftOutlinedIcon/>
                        </IconButton>
                        <Course courseId={courseId} onClick={goToTopic}/>
                    </div>
                );
            case 'topic':
                return (
                  <div>
                      <IconButton onClick={backToCourse} aria-label="edit">
                          <ArrowCircleLeftOutlinedIcon/>
                      </IconButton>
                  </div>
                );
        }
    }
    return (
        <div>
            {resolveStage()}
        </div>
    );
};

export default TeacherCourses;