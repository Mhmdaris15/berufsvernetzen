"use client";
import {
  Navbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
  NavbarMenuToggle,
  NavbarMenu,
  NavbarMenuItem,
} from "@nextui-org/navbar";
import { Link } from "@nextui-org/link";
import { Button } from "@nextui-org/button";
import ThemeSwitcher from "./ThemeSwitcher";
import {
  Dropdown,
  DropdownTrigger,
  DropdownMenu,
  DropdownItem,
} from "@nextui-org/dropdown";
import {
  AppWindow,
  ChevronDown,
  Contact2,
  TimerReset,
  User2,
  Webhook,
} from "lucide-react";
import Image from "next/image";
import Logo from "@/public/Berufsvernetzen_Icon-removebg-preview.png";


export default function NavBar() {
  const menuItems = ["jobs", "features", "dashboard", "surveys"];

  return (
    <Navbar isBlurred maxWidth="xl">
      <NavbarContent className="sm:hidden" justify="start">
        <NavbarMenuToggle />
      </NavbarContent>
      <NavbarContent className="sm:hidden pr-3" justify="center">
        <NavbarBrand>
          {/* <span className="font-bold text-inherit text-lg">Berufsvernetzen</span> */}
          <div className="w-20 h-20">
            <Image
              src={Logo}
              alt="Berufsvernetzen Logo"
              priority={true}
              width={20}
              height={20}
              className="rounded-lg"
            />

          </div>
        </NavbarBrand>
      </NavbarContent>
      <NavbarContent className="hidden sm:flex gap-5" justify="center">
        <NavbarBrand >
          <div className="flex gap-2">
            <Image
              src={Logo}
              alt="Berufsvernetzen Logo"
              priority={true}
              width={35}
              height={35}
              className="rounded-lg"
            />
            <p className="hidden font-bold text-2xl lg:block gap-3 justify-center items-center">
              Berufsvernetzen
            </p>

          </div>
        </NavbarBrand>
        <NavbarItem>
          <Button as={Link} variant="light"
            href="/user"
          >
            Jobs
          </Button>
        </NavbarItem>
        <NavbarItem>
          <Dropdown>
            <DropdownTrigger>
              <Button endContent={<ChevronDown size={16} />} variant="light">
                Features
              </Button>
            </DropdownTrigger>
            <DropdownMenu
              aria-label="ACME features"
              className="w-[340px]"
              itemClasses={{
                base: "gap-4",
              }}
            >
              <DropdownItem
                key="autoscaling"
                description="ACME scales apps to meet user demand, automagically, based on load."
                startContent={<AppWindow size={24} />}
              >
                Autoscaling
              </DropdownItem>
              <DropdownItem
                key="usage_metrics"
                description="Real-time metrics to debug issues. Slow query added? We’ll show you exactly where."
                startContent={<User2 size={24} />}
              >
                Usage Metrics
              </DropdownItem>
              <DropdownItem
                key="production_ready"
                description="ACME runs on ACME, join us and others serving requests at web scale."
                startContent={<Webhook size={24} />}
              >
                Production Ready
              </DropdownItem>
              <DropdownItem
                key="99_uptime"
                description="Applications stay on the grid with high availability and high uptime guarantees."
                startContent={<TimerReset size={24} />}
              >
                +99% Uptime
              </DropdownItem>
              <DropdownItem
                key="supreme_support"
                description="Overcome any challenge with a supporting team ready to respond."
                startContent={<Contact2 size={24} />}
              >
                +Supreme Support
              </DropdownItem>
            </DropdownMenu>
          </Dropdown>
        </NavbarItem>
        <NavbarItem>
          <Button as={Link} variant="light" href="/dashboard">
            Dashboard
          </Button>
        </NavbarItem>
        <NavbarItem>
          <Button as={Link} variant="light" href="/form/surveys">
            Surveys
          </Button>
        </NavbarItem>
      </NavbarContent>
      <NavbarContent justify="end">
        <NavbarItem className="hidden sm:flex">
          <Button
            as={Link}
            color="primary"
            href="/sign-in"
            variant="solid"
            className="hidden sm:flex"
          >
            Login
          </Button>
        </NavbarItem>
        <NavbarItem>
          <ThemeSwitcher />
        </NavbarItem>
      </NavbarContent>
      <NavbarMenu>
        {/* {menuItems.map((item, index) => (
          <NavbarMenuItem key={`${item}-${index}`}>
            <Link className="w-full" href="#" size="lg" color="foreground">
              {item}
            </Link>
          </NavbarMenuItem>
        ))} */}
        <NavbarMenuItem >
          <Link className="w-full" href="#" size="lg" color="foreground">
            Jobs
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem >
          <Link className="w-full" href="#" size="lg" color="foreground">
            Features
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem >
          <Link className="w-full" href="#" size="lg" color="foreground">
            Dashboard
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem >
          <Link className="w-full" href="/user/surveys" size="lg" color="foreground">
            Surveys
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem >
          <Link className="w-full" href="/sign-in" size="lg" color="foreground">
            Log in
          </Link>
        </NavbarMenuItem>

      </NavbarMenu>
    </Navbar>
  );
}
