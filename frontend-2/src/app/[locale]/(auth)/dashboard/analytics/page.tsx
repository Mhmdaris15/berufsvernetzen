"use client";

import React from "react";
import fs from "fs";
import path from "path";
import { useOptionAnalyticsStore } from "@/hooks/useOptionAnalyticsStore";

const DashboardAnalytics = () => {
  const { all, working, study, entrepreneurship, unemployment } =
    useOptionAnalyticsStore();

  return (
    <>
      {all && (
        <iframe
          src={`http://localhost:3000/html/survey`}
          style={{
            height: "100%",
          }}
        />
      )}
      {working && (
        <iframe
          src={`http://localhost:3000/html/survey/working`}
          style={{
            height: "100%",
          }}
        />
      )}
      {study && (
        <iframe
          src={`http://localhost:3000/html/survey/study`}
          style={{
            height: "100%",
          }}
        />
      )}
      {entrepreneurship && (
        <iframe
          src={`http://localhost:3000/html/survey/entrepreneurship`}
          style={{
            height: "100%",
          }}
        />
      )}
      {unemployment && (
        <iframe
          style={{
            height: "100%",
          }}
          src={`http://localhost:3000/html/survey/unemployment`}
        />
      )}
    </>
  );
};

export default DashboardAnalytics;
