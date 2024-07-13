'use client'
import React from "react";
import { Card, CardHeader, CardBody, Avatar, Button } from "@nextui-org/react";
import { UserIcon } from "lucide-react";

export default function App() {

    return (
        <>
            {[...Array(10)].map((_, index) => (
                <Card key={index}>
                    <CardHeader className="justify-between">
                        <div className="flex gap-5">
                            <Avatar isBordered radius="full" size="md" src="https://nextui.org/avatars/avatar-1.png" />
                            <div className="flex flex-col gap-1 items-start justify-center">
                                <h4 className="text-medium font-semibold leading-none text-default-600">Hardware Intern</h4>
                                <h5 className="text-medium tracking-tight text-default-400">Design Team Lead at Avaloq</h5>
                            </div>
                        </div>
                    </CardHeader>
                    <CardBody className="px-3 py-0 text-small text-default-400">
                        <p className="py-1">
                            <UserIcon className="inline-block" /> Internship
                        </p>
                        <p className="py-1">
                            <UserIcon className="inline-block" /> Jakarta Raya, Indonesia • On-Site
                        </p>
                        <p className="py-1">
                            <UserIcon className="inline-block" /> Rp. 1.200.000 - 1.250.000
                        </p>
                    </CardBody>
                </Card>
            ))}
        </>
    );
}
