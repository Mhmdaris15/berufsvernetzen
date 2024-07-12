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
          src={`https://berufsvernetzen.tech/html/survey`}
          style={{
            height: "100%",
          }}
        />
      )}
      {working && (
        <iframe
          src={`https://berufsvernetzen.tech/html/survey/working`}
          style={{
            height: "100%",
          }}
        />
      )}
      {study && (
        <iframe
          src={`https://berufsvernetzen.tech/html/survey/study`}
          style={{
            height: "100%",
          }}
        />
      )}
      {entrepreneurship && (
        <iframe
          src={`https://berufsvernetzen.tech/html/survey/entrepreneurship`}
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
          src={`https://berufsvernetzen.tech/html/survey/unemployment`}
        />
      )}
    </>
  );
};

export default DashboardAnalytics;
