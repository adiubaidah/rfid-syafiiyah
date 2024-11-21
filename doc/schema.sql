-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2024-11-21T16:31:01.009Z

CREATE TYPE "role_type" AS ENUM (
  'superadmin',
  'admin',
  'employee',
  'parent'
);

CREATE TYPE "gender_type" AS ENUM (
  'male',
  'female'
);

CREATE TYPE "presence_type" AS ENUM (
  'alpha',
  'permission',
  'sick',
  'late',
  'present'
);

CREATE TYPE "presence_created_by_type" AS ENUM (
  'system',
  'tap',
  'admin'
);

CREATE TYPE "santri_permission_type" AS ENUM (
  'sick',
  'permission'
);

CREATE TYPE "device_mode_type" AS ENUM (
  'record',
  'presence',
  'permission',
  'ping'
);

CREATE TABLE "user" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "role" role_type,
  "username" varchar(50) UNIQUE,
  "password" varchar(255)
);

CREATE TABLE "holiday" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "color" char(7),
  "description" varchar(255)
);

CREATE TABLE "holiday_date" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "date" date NOT NULL,
  "holiday_id" int NOT NULL
);

CREATE TABLE "santri_schedule" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" varchar(255),
  "start_presence" time NOT NULL,
  "start_time" time NOT NULL,
  "finish_time" time NOT NULL
);

CREATE TABLE "santri_occupation" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "description" varchar(255)
);

CREATE TABLE "santri" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "nis" varchar(15) UNIQUE,
  "name" varchar(255) NOT NULL,
  "gender" gender_type NOT NULL,
  "generation" int NOT NULL,
  "is_active" boolean DEFAULT true,
  "photo" varchar(100),
  "occupation_id" int,
  "parent_id" int
);

CREATE TABLE "parent" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "address" varchar(255) NOT NULL,
  "gender" gender_type NOT NULL,
  "whatsapp_number" varchar(14) UNIQUE,
  "photo" varchar(100),
  "user_id" int UNIQUE
);

CREATE TABLE "smart_card" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "uid" varchar(20) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "is_active" boolean NOT NULL DEFAULT false,
  "santri_id" int,
  "employee_id" int
);

CREATE TABLE "santri_presence" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY,
  "schedule_id" int NOT NULL,
  "schedule_name" varchar(100) NOT NULL,
  "type" presence_type NOT NULL,
  "santri_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" presence_created_by_type NOT NULL,
  "notes" text,
  "santri_permission_id" int
);

CREATE TABLE "santri_permission" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "santri_id" int NOT NULL,
  "type" santri_permission_type NOT NULL,
  "start_permission" timestamptz NOT NULL DEFAULT 'now()',
  "end_permission" timestamptz,
  "excuse" varchar(255) NOT NULL
);

CREATE TABLE "employee_occupation" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" varchar(255)
);

CREATE TABLE "employee" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "nip" char(18) UNIQUE,
  "name" varchar(100) NOT NULL,
  "gender" gender_type NOT NULL,
  "photo" varchar(100),
  "occupation_id" int NOT NULL,
  "user_id" int UNIQUE
);

CREATE TABLE "admin_restrictions" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "admin_id" int NOT NULL,
  "restricted_employee_id" int NOT NULL
);

CREATE TABLE "employee_schedule" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "start_presence" time NOT NULL,
  "start_time" time NOT NULL,
  "finish_time" time NOT NULL
);

CREATE TABLE "employee_presence" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY,
  "schedule_id" int,
  "type" presence_type NOT NULL,
  "employee_id" int NOT NULL,
  "notes" text
);

CREATE TABLE "employee_permission" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "employee_id" int NOT NULL,
  "schedule_id" int NOT NULL,
  "schedule_name" varchar(100) NOT NULL,
  "start_permission" time NOT NULL DEFAULT 'now()',
  "end_permission" time,
  "reason" varchar(255) NOT NULL,
  "is_go_home" boolean
);

CREATE TABLE "device" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL
);

CREATE TABLE "device_mode" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "mode" device_mode_type NOT NULL,
  "input_topic" varchar(100) NOT NULL,
  "acknowledgment_topic" varchar(100) NOT NULL,
  "device_id" int NOT NULL
);

CREATE INDEX ON "user" ("username");

CREATE UNIQUE INDEX ON "santri_schedule" ("start_presence", "start_time", "finish_time");

CREATE INDEX ON "employee" ("name");

CREATE UNIQUE INDEX ON "admin_restrictions" ("admin_id", "restricted_employee_id");

COMMENT ON COLUMN "holiday"."name" IS 'Optional description of the holiday';

COMMENT ON COLUMN "santri_schedule"."start_time" IS 'Waktu mulai kegiatan';

COMMENT ON COLUMN "santri_schedule"."finish_time" IS 'Waktu berakhirnya kegiatan';

COMMENT ON COLUMN "santri"."generation" IS 'ex: 2024, 2022';

COMMENT ON COLUMN "smart_card"."santri_id" IS 'Smart Card bisa milik santri';

COMMENT ON COLUMN "smart_card"."employee_id" IS 'Smart Card bisa milik employee';

COMMENT ON COLUMN "santri_presence"."schedule_id" IS 'Karena bisa saja activitynya dihapus';

COMMENT ON COLUMN "santri_presence"."schedule_name" IS 'menggunakan name, karena jika activity dihapus, atau diubah maka masih tetap ada presence nya, karena bersifat history';

COMMENT ON COLUMN "santri_presence"."santri_permission_id" IS 'Jika izin ditengah kegiatan maka akan diisi';

COMMENT ON COLUMN "santri_permission"."end_permission" IS 'Waktu berakhir, jika pulang, maka setting end permissionnya di akhir waktu berakhirnya schedule yang terakhir';

COMMENT ON COLUMN "employee_schedule"."name" IS 'ex: Pagi, siang, sore, malam';

COMMENT ON COLUMN "employee_schedule"."start_time" IS 'Waktu jenis';

COMMENT ON COLUMN "employee_permission"."end_permission" IS 'waktu kembali, null berarti pulang';

COMMENT ON COLUMN "employee_permission"."is_go_home" IS 'Pulang, keluar sementara';

COMMENT ON COLUMN "device"."name" IS 'ex: device1';

COMMENT ON COLUMN "device_mode"."input_topic" IS 'topic for input';

COMMENT ON COLUMN "device_mode"."acknowledgment_topic" IS 'topic for acknowledgment';

ALTER TABLE "holiday_date" ADD FOREIGN KEY ("holiday_id") REFERENCES "holiday" ("id") ON DELETE CASCADE;

ALTER TABLE "santri" ADD FOREIGN KEY ("occupation_id") REFERENCES "santri_occupation" ("id") ON DELETE SET NULL;

ALTER TABLE "santri" ADD FOREIGN KEY ("parent_id") REFERENCES "parent" ("id") ON DELETE SET NULL;

ALTER TABLE "parent" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE SET NULL;

ALTER TABLE "smart_card" ADD FOREIGN KEY ("santri_id") REFERENCES "santri" ("id");

ALTER TABLE "smart_card" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("id");

ALTER TABLE "santri_presence" ADD FOREIGN KEY ("santri_id") REFERENCES "santri" ("id") ON DELETE CASCADE;

ALTER TABLE "santri_presence" ADD FOREIGN KEY ("santri_permission_id") REFERENCES "santri_permission" ("id") ON DELETE CASCADE;

ALTER TABLE "santri_permission" ADD FOREIGN KEY ("santri_id") REFERENCES "santri" ("id");

ALTER TABLE "employee" ADD FOREIGN KEY ("occupation_id") REFERENCES "employee_occupation" ("id");

ALTER TABLE "employee" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE SET NULL;

ALTER TABLE "admin_restrictions" ADD FOREIGN KEY ("admin_id") REFERENCES "employee" ("id");

ALTER TABLE "admin_restrictions" ADD FOREIGN KEY ("restricted_employee_id") REFERENCES "employee" ("id");

ALTER TABLE "employee_presence" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("id") ON DELETE CASCADE;

ALTER TABLE "employee_permission" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("id") ON DELETE CASCADE;

ALTER TABLE "device_mode" ADD FOREIGN KEY ("device_id") REFERENCES "device" ("id") ON DELETE CASCADE;
