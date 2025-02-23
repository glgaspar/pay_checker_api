create database paychecker;

create table bills (
    id          serial primary key,           --id is just id
    description varchar(50) not null,         --name that bill
    expDay      int not null,                 --expiration day
    lastDate    timestamp null,               --last payment date
    path        varchar(25) not null,         --directory for receipts
    track       boolean not null default true --track payment
)
