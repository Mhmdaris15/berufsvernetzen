"use client";

import React, { useEffect, useState } from 'react';
import ReactMarkdown from 'react-markdown';

type Feedback = string;

interface Props {}

const AIAnalyzerPage: React.FC<Props> = () => {
  const [feedbacks, setFeedbacks] = useState<Feedback[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchFeedbacks();
  }, []);

  const fetchFeedbacks = async () => {
    try {
      setIsLoading(true);
      // Replace this with your actual API endpoint
      const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/surveys/feedbacks`);
      if (!response.ok) {
        console.log(response);
        throw new Error('Failed to fetch feedbacks');
      }
      const data = await response.json();
      setFeedbacks(data.data);
    } catch (err) {
      console.error(err);
      setError('Failed to load feedbacks. Please try again later.');
    } finally {
      setIsLoading(false);
    }
  };

  const renderFeedback = (feedback: Feedback, index: number) => (
    <div key={index} className="mb-8 p-4 border rounded shadow">
      <ReactMarkdown>{feedback}</ReactMarkdown>
    </div>
  );

  if (isLoading) {
    return <div className="text-center mt-8">Loading...</div>;
  }

  if (error) {
    return <div className="text-center mt-8 text-red-500">{error}</div>;
  }

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">AI Feedback Analyzer</h1>
      {feedbacks.length > 0 ? (
        feedbacks.map(renderFeedback)
      ) : (
        <p>No feedbacks available.</p>
      )}
    </div>
  );
};

export default AIAnalyzerPage;