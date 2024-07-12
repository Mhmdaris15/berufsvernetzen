"use client";

import React from 'react'
import ArisPhoto from "@/public/Profile Picture Aris.jpg";
import Kebersamaan from "@/public/kebersaamaan.jpg";
import Kebersamaan2 from "@/public/kebersamaan 2.jpg";
import Kebersamaan3 from "@/public/kebersamaan 3.jpg";
import Image from 'next/image';
import Link from 'next/link';
import { useTheme } from 'next-themes';


type Props = {}

const Devs = (props: Props) => {
  const { theme } = useTheme();

  return (
    <>
      <h1 className={`text-3xl font-bold text-center ${theme == "dark" ? "text-white" : "text-gray-800"}`}>
        Meet The Developers!
      </h1>
      <div className="flex flex-col md:flex-row md:flex-wrap justify-center gap-5 px-7 md:gap-10 mt-10">
        <Link
          href=""
          target="_blank"
          rel="noopener noreferrer"
          className={`flex flex-col items-center p-6 rounded-lg min-w-64 ${theme == "dark"
            ? "bg-blue-400 bg-opacity-20"
            : "bg-green-100 shadow-md"
            } transform hover:scale-105 transition-transform`}
        >
          <Image
            src={ArisPhoto}
            alt="Leader Developer"
            width={80}
            height={80}
            className="rounded-full mb-4"
          />
          <h3 className={`font-semibold ${theme == "dark" ? "text-white" : "text-gray-800"}`}>Aris Septanugroho</h3>
          <p className={` ${theme == "dark" ? "text-white" : "text-gray-600"}`}>Machine Learning Engineer</p>
          <p className={` ${theme == "dark" ? "text-white" : "text-gray-600"}`}>PT Matra Hillindo</p>
        </Link>
        <Link
          href="#"
          target="_blank"
          rel="noopener noreferrer"
          className={`flex flex-col items-center p-6 rounded-lg min-w-64 ${theme == "dark"
            ? "bg-white bg-opacity-20"
            : "bg-yellow-200 shadow-md"
            } transform hover:scale-105 transition-transform`}
        >
          <Image
            src={Kebersamaan3}
            alt="Developer 3"
            width={80}
            height={80}
            className="rounded-full mb-4"
          />
          <h3 className={`font-semibold ${theme == "dark" ? "text-white" : "text-gray-800"}`}>Fadli</h3>
          <p className={` ${theme == "dark" ? "text-white" : "text-gray-600"}`}>Front End Engineer</p>
          <p className={` ${theme == "dark" ? "text-white" : "text-gray-600"}`}>PT Matra Hillindo</p>
        </Link>
        <Link
          href=""
          target="_blank"
          rel="noopener noreferrer"
          className={`flex flex-col items-center p-6 rounded-lg min-w-64 ${theme == "dark"
            ? "bg-red-400 bg-opacity-20"
            : "bg-blue-100 shadow-md"
            } transform hover:scale-105 transition-transform`}
        >
          <Image
            src={Kebersamaan}
            alt="Developer 1"
            width={80}
            height={80}
            className="rounded-full mb-4"
          />
          <h3 className={`font-semibold ${theme == "dark" ? "text-white" : "text-gray-800"}`}>Farhan</h3>
          <p className={` ${theme == "dark" ? "text-white" : "text-gray-600"}`}>UI/UX Designer</p>
          <p className={` ${theme == "dark" ? "text-white" : "text-gray-600"}`}>PT Matra Hillindo</p>
        </Link>




        <Link
          href="#"
          target="_blank"
          rel="noopener noreferrer"
          className={`flex flex-col items-center p-6 rounded-lg min-w-64 ${theme == "dark"
            ? "bg-white bg-opacity-20"
            : "bg-yellow-200 shadow-md"
            } transform hover:scale-105 transition-transform`}
        >
          <Image
            src={Kebersamaan2}
            alt="Developer 2"
            width={80}
            height={80}
            className="rounded-full mb-4"
          />
          <h3 className={`font-semibold ${theme == "dark" ? "text-white" : "text-gray-800"}`}>Bryant</h3>
          <p className={` ${theme == "dark" ? "text-white" : "text-gray-600"}`}>FrontEnd Engineer</p>
          <p className={` ${theme == "dark" ? "text-white" : "text-gray-600"}`}>PT Matra Hillindo</p>
        </Link>
      </div>
    </>
  );
}

export default Devs