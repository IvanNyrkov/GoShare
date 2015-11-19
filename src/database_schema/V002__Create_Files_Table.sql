SET search_path = model, pg_catalog;

CREATE TABLE uploaded_file (
  uploaded_file_id      INTEGER NOT NULL,
  path                  CHARACTER VARYING(512) NOT NULL,
  passphrase            CHARACTER VARYING(512) NOT NULL,
  created_at            TIMESTAMP
);
ALTER TABLE model.uploaded_file OWNER TO admin;

CREATE SEQUENCE uploaded_file_uploaded_file_id_seq
START WITH 1
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;
ALTER TABLE model.uploaded_file_uploaded_file_id_seq OWNER TO admin;
ALTER SEQUENCE uploaded_file_uploaded_file_id_seq OWNED BY uploaded_file.uploaded_file_id;
ALTER TABLE ONLY uploaded_file ALTER COLUMN uploaded_file_id SET DEFAULT nextval('uploaded_file_uploaded_file_id_seq' :: REGCLASS);

ALTER TABLE ONLY uploaded_file
ADD CONSTRAINT uploaded_file_pkey PRIMARY KEY (uploaded_file_id);
