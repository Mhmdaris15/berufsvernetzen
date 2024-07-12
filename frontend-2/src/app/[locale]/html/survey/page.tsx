import React from 'react'
import fs from 'fs';
import path from 'path';

type Props = {}

const SurveyDataProfilingPage = (props: Props) => {
  const htmlContent = fs.readFileSync(path.join(process.cwd(), 'public/html/entrepreneur_survey_data_profiling.html'), 'utf-8');
  return (
    <div dangerouslySetInnerHTML={{ __html: htmlContent }} />
    // <iframe src="html/entrepreneur_survey_data_profiling.html" frameborder="0"></iframe>
  )
}

export default SurveyDataProfilingPage