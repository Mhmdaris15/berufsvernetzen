"use client";
import { Accordion, AccordionItem } from "@nextui-org/accordion";
import { motion } from "framer-motion";

const faqs = [
  {
    question: "What is Berufsvernetzen?",
    answer: "Berufsvernetzen is a platform designed to connect alumni of a school, facilitating networking and career opportunities."
  },
  {
    question: "Who can use Berufsvernetzen?",
    answer: "Berufsvernetzen is open to alumni, current students, school staff, and other stakeholders."
  },
  // {
  //   question: "Is Berufsvernetzen free to use?",
  //   answer: "Yes, Berufsvernetzen is completely free for all users."
  // },
  {
    question: "How do I Log In?",
    answer: "You can log in after your school collaborates with us!."
  },
  {
    question: "How does Berufsvernetzen protect my privacy?",
    answer: "We prioritize the privacy and security of our users' personal information and have implemented measures to safeguard it."
  },
  {
    question: "Can Alumni search for job opportunities on Berufsvernetzen?",
    answer: "Yes, Alumni can explore job listings and opportunities available through Berufsvernetzen."
  },
  // {
  //   question: "Can I connect with other alumni on Berufsvernetzen?",
  //   answer: "Absolutely! Berufsvernetzen provides networking features such as messaging, friend requests, and alumni directories to help you connect with fellow alumni."
  // },
  {
    question: "How can my school benefit from Berufsvernetzen?",
    answer: "Schools can benefit from Berufsvernetzen by keeping track of alumni data, fostering alumni engagement, and accessing analytics and insights about their alumni network."
  },
  // {
  //   question: "What types of events are hosted on Berufsvernetzen?",
  //   answer: "Berufsvernetzen hosts events, webinars, and workshops related to career development and industry insights. You can participate or access recordings through the platform."
  // },
  // {
  //   question: "How can I provide feedback or suggest new features?",
  //   answer: "We welcome your feedback and feature requests! You can reach out to our support team through the provided channels."
  // },
  {
    question: "Is Berufsvernetzen accessible on mobile devices?",
    answer: "Yes, Berufsvernetzen is optimized for mobile use, allowing you to access it conveniently from your smartphone or tablet."
  },
  // {
  //   question: "How can I delete my account?",
  //   answer: "If you wish to delete your Berufsvernetzen account, you can follow the account deletion process outlined on our website."
  // }
];


export default function Faq() {

  return (
    <section className="relative max-w-screen-2xl mx-auto px-4 py-28 gap-12 md:px-8 flex flex-col justify-center items-center">
      <motion.div
        initial={{ y: 5, opacity: 0 }}
        whileInView={{
          y: 0,
          opacity: 1,
        }}
        viewport={{ once: true }}
        transition={{ duration: 1, delay: 0.5 }}
        className="flex flex-col gap-3 justify-center items-center"
      >
        <h4 className="text-2xl font-bold sm:text-3xl bg-gradient-to-b from-foreground to-foreground/70 text-transparent bg-clip-text">
          FAQ
        </h4>
        <p className="max-w-xl text-foreground/80 text-center">
          Here are some of our frequently asked questions. If you have any other
          questions you’d like answered please feel free to email us.
        </p>
      </motion.div>
      <motion.div
        initial={{ y: 5, opacity: 0 }}
        whileInView={{
          y: 0,
          opacity: 1,
        }}
        viewport={{ once: true }}
        transition={{ duration: 1, delay: 1 }}
        className="max-w-2xl w-full border border-foreground/50 rounded-md p-1"
      >
        <Accordion
          motionProps={{
            variants: {
              enter: {
                y: 0,
                opacity: 1,
                height: "auto",
                transition: {
                  height: {
                    type: "spring",
                    stiffness: 500,
                    damping: 30,
                    duration: 1,
                  },
                  opacity: {
                    easings: "ease",
                    duration: 1,
                  },
                },
              },
              exit: {
                y: -10,
                opacity: 0,
                height: 0,
                transition: {
                  height: {
                    easings: "ease",
                    duration: 0.25,
                  },
                  opacity: {
                    easings: "ease",
                    duration: 0.3,
                  },
                },
              },
            },
          }}
        >
          {faqs.map((faq, index) => (
            <AccordionItem key={index} title={faq.question}>
              <p>{faq.answer}</p>
            </AccordionItem>
          ))}
        </Accordion>
      </motion.div>
      <div className="absolute w-40 h-40 border bg-light -left-28 top-0 blur-[200px]"></div>
    </section>
  );
}
