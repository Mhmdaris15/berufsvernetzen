"use client";

import React from 'react'
import { json } from "@/types/surveys"
import { Survey } from "survey-react-ui"
import { Model } from "survey-core";
import { ContrastDark } from "survey-core/themes/contrast-dark";
import "survey-core/defaultV2.min.css";


const SurveysFormPage = () => {
  const survey = new Model(json);
  const [pageNo, setPageNo] = React.useState<number>(0);

  React.useEffect(() => {
    survey.applyTheme(ContrastDark);
    
    setPageNo(survey.currentPageNo);
  }, [survey]); // Apply theme only when survey object changes 

  survey.onComplete.add((sender, options) => {
    console.log(JSON.stringify(sender.data, null, 3));
});

  return (
    <Survey currentPageNo={pageNo} model={survey} />
  )
}

export default SurveysFormPage