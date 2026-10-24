# User:

- id : int
- name : varchar
- occupation : varchar
- email : varchar
- password_hash : varchar
- avatar_file_name : varchar
- role : varchar
- token : varchar
- created_at : datetime
- updated_at : datetime

# Campaign :

- id : int
- user_id : int
- title : varchar
- short_description : varchar
- description : text
- goal_amount : int
- current_amount : int
- perks : text
- backer_count (org yg mendanai): int
- slug : varchar
- created_at : datetime
- updated_at : datetime

# Campaign Image :

- id : int
- campaign_id: int
- file_name : varchar
- is_primary : boolan (tinyint)
- created_at : datetime
- updated_at : datetime

# Transaction:

- id : int
- campaign_id : int
- user_id : int
- amount : int
- status : varchar
- code : varchar (kode transaksi)
- created_at : datetime
- updated_at : datetime
