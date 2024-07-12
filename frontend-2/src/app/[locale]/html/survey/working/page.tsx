import React from 'react'
import fs from 'fs';
import path from 'path';

type Props = {}

const WorkingSurveyDataProfilingPage = (props: Props) => {
  const htmlContent = fs.readFileSync(path.join(process.cwd(), 'public/html/working_survey_data_profiling.html'), 'utf-8');
  return (
    <div dangerouslySetInnerHTML={{ __html: htmlContent }} />
    // <iframe src="html/entrepreneur_survey_data_profiling.html" frameborder="0"></iframe>
  )
}

export default WorkingSurveyDataProfilingPage