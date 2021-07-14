-- Table: public.devtech

-- DROP TABLE public.devtech;

CREATE TABLE public.devtech
(
    id integer NOT NULL DEFAULT nextval('"devTech_id_seq"'::regclass),
    rank integer,
    name text COLLATE pg_catalog."default",
    CONSTRAINT "devTech_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.devtech
    OWNER to postgres;