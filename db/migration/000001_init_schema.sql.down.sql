-- Drop tables if they exist

DROP TABLE IF EXISTS "user" CASCADE;
DROP TABLE IF EXISTS "holiday_day" CASCADE;
DROP TABLE IF EXISTS "holiday" CASCADE;
DROP TABLE IF EXISTS "santri_schedule" CASCADE;
DROP TABLE IF EXISTS "santri_occupation" CASCADE;
DROP TABLE IF EXISTS "santri" CASCADE;
DROP TABLE IF EXISTS "parent" CASCADE;
DROP TABLE IF EXISTS "rfid" CASCADE;
DROP TABLE IF EXISTS "santri_presence" CASCADE;
DROP TABLE IF EXISTS "santri_permission" CASCADE;
DROP TABLE IF EXISTS "employee_occupation" CASCADE;
DROP TABLE IF EXISTS "employee" CASCADE;
DROP TABLE IF EXISTS "employee_schedule" CASCADE;
DROP TABLE IF EXISTS "employee_presence" CASCADE;
DROP TABLE IF EXISTS "employee_permission" CASCADE;
DROP TABLE IF EXISTS "arduino" CASCADE;
DROP TABLE IF EXISTS "arduino_mode" CASCADE;

-- Drop enums if they exist
DROP TYPE IF EXISTS "user_role" CASCADE;
DROP TYPE IF EXISTS "gender" CASCADE;
DROP TYPE IF EXISTS "presence_type" CASCADE;
DROP TYPE IF EXISTS "arduino_mode_type" CASCADE;
DROP TYPE IF EXISTS "santri_permission_type" CASCADE;
DROP TYPE IF EXISTS "presence_created_by" CASCADE;