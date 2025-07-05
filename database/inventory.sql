--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: contacts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contacts (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    message text NOT NULL
);


ALTER TABLE public.contacts OWNER TO postgres;

--
-- Name: contacts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contacts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.contacts_id_seq OWNER TO postgres;

--
-- Name: contacts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contacts_id_seq OWNED BY public.contacts.id;


--
-- Name: experiences; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.experiences (
    id integer NOT NULL,
    company character varying(100) NOT NULL,
    role character varying(100) NOT NULL,
    description text,
    start_date date,
    end_date date
);


ALTER TABLE public.experiences OWNER TO postgres;

--
-- Name: experiences_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.experiences_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.experiences_id_seq OWNER TO postgres;

--
-- Name: experiences_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.experiences_id_seq OWNED BY public.experiences.id;


--
-- Name: projects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.projects (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    description text NOT NULL,
    tech_stack character varying(100),
    start_date date,
    end_date date
);


ALTER TABLE public.projects OWNER TO postgres;

--
-- Name: projects_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.projects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.projects_id_seq OWNER TO postgres;

--
-- Name: projects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.projects_id_seq OWNED BY public.projects.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    headline character varying(150),
    bio text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: contacts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contacts ALTER COLUMN id SET DEFAULT nextval('public.contacts_id_seq'::regclass);


--
-- Name: experiences id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiences ALTER COLUMN id SET DEFAULT nextval('public.experiences_id_seq'::regclass);


--
-- Name: projects id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects ALTER COLUMN id SET DEFAULT nextval('public.projects_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: contacts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contacts (id, name, email, message) FROM stdin;
1	Zahrazahra Putri	zahrazahraputri57@gmail.com	test website portofolio zahra
2	zahra	admin@gmail.com	tes 2
3	zahraa	admin@example.com	test 3
\.


--
-- Data for Name: experiences; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.experiences (id, company, role, description, start_date, end_date) FROM stdin;
1	Universitas Ahmad Dahlan - FTI	Asisten Praktikum	• Membimbing mahasiswa dalam memahami konsep-konsep utama di bidang Machine Learning, Statistika Informatika, dan Sistem Pendukung Keputusan\n• Memberikan bimbingan teknis terkait analisis data, pemrograman, atau topik sesuai mata kuliah.\n• Membantu menyelesaikan permasalahan praktis dan memastikan kelancaran pelaksanaan praktikum.\n• Mendapatkan penghargaan atas dedikasi dan kinerja yang konsisten selama menjadi asisten.	2023-10-28	2024-01-07
\.


--
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.projects (id, title, description, tech_stack, start_date, end_date) FROM stdin;
1	To-Do List CLI Scripting	Aplikasi sederhana menggunakan bash script di Ubuntu untuk membantu pengguna mengelola daftar tugas melalui terminal.	Bash	2025-08-05	2025-08-05
2	To-Do List CLI	Aplikasi berbasis CLI menggunakan Golang yang bertujuan untuk membantu pengguna mengelola daftar tugas mereka secara lokal melalui terminal, dengan menyimpan data ke dalam file berformat JSON. Aplikasi ini menyediakan fitur untuk menambahkan, menampilkan, menyelesaikan, dan menghapus tugas.	Golang	2025-05-25	2025-05-26
3	Sistem Inventaris Barang Kantor	Aplikasi sistem inventaris kantor berbasis Command Line Interface (CLI) menggunakan bahasa pemrograman Golang. Aplikasi ini bertujuan untuk membantu pengguna mengelola data kategori dan barang inventaris melalui terminal, dengan menyimpan data pada database.	Golang , PostgreSQL	2025-06-15	2025-06-16
4	FreshEats	Model klasifikasi image bahan makanan menggunakan algoritma machine learning dengan arsitektur MobileNet.	Python , Scikit-learn , Pandas	2023-12-14	2023-12-21
5	Grevee	Hasil analisis menggunakan data yang diberikan divisi backend yang disajikan untuk stakeholder menggunakan Looker Studio.	Python , Scikit-learn , Pandas	2024-05-25	2024-05-28
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, headline, bio) FROM stdin;
1	Zahra Putri Zanuarti	zahraputrizanuarti01@gmail.com	Data Engineer | Golang & Python Enthusiast	Lulusan Informatika dengan fokus pada Data Engineering dan pengembangan backend menggunakan Golang dan Python. Berpengalaman dalam ETL, analisis data, dan integrasi sistem backend berbasis web. Alumni Bangkit Academy (ML Cohort) dan mentee Data Engineer di Alterra Academy. Antusias membangun solusi data end-to-end yang scalable dan efisien.
\.


--
-- Name: contacts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.contacts_id_seq', 3, true);


--
-- Name: experiences_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.experiences_id_seq', 2, true);


--
-- Name: projects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.projects_id_seq', 5, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: contacts contacts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contacts
    ADD CONSTRAINT contacts_pkey PRIMARY KEY (id);


--
-- Name: experiences experiences_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiences
    ADD CONSTRAINT experiences_pkey PRIMARY KEY (id);


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

