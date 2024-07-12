import { auth } from '@clerk/nextjs/server';
import { redirect } from 'next/navigation';
import Logo from "@/public/Berufsvernetzen_Icon-removebg-preview.png";
import Image from 'next/image';


export default function CenteredLayout(props: { children: React.ReactNode }) {
  const { userId } = auth();

  if (userId) {
    redirect('/dashboard');
  }

  return (
    <div className="mx-auto py-8 px-6">
      <div>
        <div className='text-center'>
          <Image
            src={Logo}
            alt="Berufsvernetzen Logo"
            priority={true}
            width={130}
            height={130}
            className="rounded-lg mx-auto"
          />
          <h1 className="text-2xl font-extrabold mx-auto lg:text-4xl">
            Berufsvernetzen
          </h1>
        </div>
        <div className='flex justify-center'>
          {props.children}
        </div>
      </div>
    </div>
  );
}
