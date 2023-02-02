CREATE TABLE "book" (
  "id" bigserial PRIMARY KEY,
  "isbn" int,
  "title" varchar,
  "author" varchar,
  "publication_date" timestamptz,
  "publisher" varchar,
  "genre" varchar,
  "language" varchar,
  "format" enum,
  "available_copies" int,
  "total_copies" int,
  "description" varchar,
  "cover_image" varchar
);

CREATE TABLE "member" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "email" varchar,
  "student_id" int,
  "staff_id" int,
  "membership_date" timestamptz DEFAULT (now()),
  "address" varchar,
  "phone_number" int,
  "membership_type" enum,
  "library_card_number" int
);

CREATE TABLE "borrow" (
  "id" bigserial PRIMARY KEY,
  "member_id" bigserial,
  "book_id" bigserial,
  "borrow_date" timestamptz DEFAULT (now()),
  "due_date" timestamptz,
  "return_date" timestamptz,
  "status" enum
);

CREATE TABLE "return" (
  "id" bigserial PRIMARY KEY,
  "borrow_id" bigserial,
  "return_date" timestamptz DEFAULT (now()),
  "late_fee" int,
  "returned_by_id" bigserial
);

CREATE TABLE "reservation" (
  "id" bigserial PRIMARY KEY,
  "member_id" bigserial,
  "book_id" bigserial,
  "reservation_date" timestamptz DEFAULT (now()),
  "reservation_status" enum,
  "fulfill_date" timestamptz
);

CREATE TABLE "overdueNotice" (
  "id" bigserial PRIMARY KEY,
  "borrow_id" bigserial,
  "notice_date" timestamptz,
  "noticed_by" bigserial,
  "notice_type" enum
);

CREATE TABLE "fines" (
  "id" bigserial PRIMARY KEY,
  "member_id" bigserial,
  "amount" int,
  "description" varchar,
  "date_assessed" timestamptz,
  "date_paid" timestamptz,
  "payment_method" enum
);

CREATE TABLE "review" (
  "id" bigserial PRIMARY KEY,
  "book_id" bigserial,
  "member_id" bigserial,
  "rating" int,
  "review_text" varchar,
  "review_date" timestamptz
);

CREATE TABLE "recommendation" (
  "id" bigserial,
  "book_id" bigserial,
  "member_id" bigserial,
  "recommendation_text" varchar,
  "recommendation_date" timestamptz,
  "recommended_to" varchar
);

ALTER TABLE "borrow" ADD FOREIGN KEY ("member_id") REFERENCES "member" ("id");

ALTER TABLE "borrow" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("id");

ALTER TABLE "return" ADD FOREIGN KEY ("borrow_id") REFERENCES "borrow" ("id");

ALTER TABLE "return" ADD FOREIGN KEY ("returned_by_id") REFERENCES "member" ("id");

ALTER TABLE "reservation" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("id");

ALTER TABLE "reservation" ADD FOREIGN KEY ("member_id") REFERENCES "member" ("id");

ALTER TABLE "overdueNotice" ADD FOREIGN KEY ("borrow_id") REFERENCES "borrow" ("id");

ALTER TABLE "overdueNotice" ADD FOREIGN KEY ("noticed_by") REFERENCES "member" ("id");

ALTER TABLE "fines" ADD FOREIGN KEY ("member_id") REFERENCES "member" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("member_id") REFERENCES "member" ("id");

ALTER TABLE "recommendation" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("id");

ALTER TABLE "recommendation" ADD FOREIGN KEY ("member_id") REFERENCES "member" ("id");
