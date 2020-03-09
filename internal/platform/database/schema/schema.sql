create database traffic_violation;
create type gender as enum('L', 'P');
create type marriage_status as enum('single', 'married', 'widow', 'widower');
create type religion as enum('islam', 'kristen', 'budha', 'hindu', 'kong hu cu');
create type citizen_address_type as enum('primary', 'secondary');
create type vehicle_type as enum('motor', 'mobil');

create table citizens (
  id serial primary key,
  uuid varchar(36) not null,
  nik varchar(16) unique not null,
  name varchar(100) not null,
  gender gender not null,
  phone_number varchar(15) not null,
  email varchar(100),
  status marriage_status not null,
  place_of_birth varchar(100),
  date_of_birth date not null,
  religion religion not null,
  created_at timestamptz not null,
  updated_at timestamptz not null,
  deleted_at timestamptz null
);

create table citizen_addresses (
   id serial primary key,
   citizen_id int8 not null,
   address_line_1 text,
   address_line_2 text,
   address_type citizen_address_type not null default 'primary',
   province varchar(100),
   city varchar(100),
   subdistrict varchar(100),
   rt varchar(5),
   rw varchar(5),
   zipcode varchar(5),
   created_at timestamptz not null,
   updated_at timestamptz not null,
   deleted_at timestamptz null,
   FOREIGN KEY (citizen_id) references citizens(id)
);

create table stnk (
  id serial primary key,
  uuid varchar(36) not null,
  code varchar(100) not null,
  police_number varchar(100) not null,
  owner_name varchar(100) not null,
  adderss text,
  stnk_type varchar(100) not null,
  model vehicle_type not null,
  assembled_at date,
  color varchar(50),
  centimeter_cubic varchar(50),
  chasis_code varchar(50),
  machine_code varchar(100),
  created_at timestamptz not null,
  updated_at timestamptz not null,
  deleted_at timestamptz null
);

create table bpkb (
  id serial primary key,
  uuid varchar(36),
  code varchar(100),
  bpkb_type vehicle_type,
  police_number varchar(100),
  created_year varchar(4),
  assembled_at date,
  color varchar(50),
  centimeter_cubic varchar(50),
  created_at timestamptz not null,
  updated_at timestamptz not null,
  deleted_at timestamptz null
);

create table citizen_vehicles (
  id serial primary key,
  uuid varchar(36) not null,
  citizen_id int8 not null,
  stnk_code varchar(100) not null,
  bpkb_code varchar(100) not null,
  vehicle_type vehicle_type not null,
  created_at timestamptz not null,
  updated_at timestamptz not null,
  deleted_at timestamptz null,
  foreign key (citizen_id) references citizens(id),
  foreign key (stnk_code) references stnk(code),
  foreign key (bpkb_code) references bpkb(code)
);

create table violation_laws (
    id serial primary key,
    uuid varchar(36) not null,
    law varchar(100) not null,
    content text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz null
);

create table vehicle_violations (
    id serial primary key,
    uuid varchar(36) not null,
    violation_id int8 not null,
    dishub_office_id int8 not null,
    police_number varchar(100) not null,
    latitude double precision not null,
    longiture double precision not null,
    violated_at date not null,
    processed bool not null default false,
    processed_at date,
    evidence_img_url text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz null
);

create table dishub_offices (
    id serial primary key,
    uuid varchar(36) not null,
    email varchar(100) not null,
    phone varchar(16) not null,
    adddress text not null,
    province varchar(100) not null,
    rt varchar(5),
    rw varchar(5),
    zipcode varchar(5),
    latitude double precision not null,
    longitude double precision not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz null
);
