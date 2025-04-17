CREATE DATABASE marifat_db;
-- Table: Users
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(100) NOT NULL,
    usersurname VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    user_number VARCHAR(9) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    user_role VARCHAR(10) DEFAULT 'user', -- Default role is 'user'
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Table: Teachers
CREATE TABLE teachers (
    teacher_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    teacher_name VARCHAR(100) NOT NULL,
    teacher_surname VARCHAR(100) NOT NULL,
    teacher_number VARCHAR(9) NOT NULL UNIQUE,
    teacher_tg VARCHAR(255) NOT NULL UNIQUE,
    bio TEXT,
    rating FLOAT DEFAULT 0,
    user_id UUID, -- Add user_id column for the foreign key reference
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);



-- Table: Courses
CREATE TABLE courses (
    course_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    teacher_id UUID NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (teacher_id) REFERENCES teachers(teacher_id) ON DELETE CASCADE
);

-- Table: Videos
CREATE TABLE videos (
    video_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL,
    teacher_id UUID NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    video_url VARCHAR(255) NOT NULL,
    upload_date TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (course_id) REFERENCES courses(course_id) ON DELETE CASCADE,
    FOREIGN KEY (teacher_id) REFERENCES teachers(teacher_id) ON DELETE CASCADE

);

-- Table: Comments
CREATE TABLE comments (
    comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    video_id UUID NOT NULL,
    user_id UUID NOT NULL,
    comment_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (video_id) REFERENCES videos(video_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Table: View Counts
CREATE TABLE view_counts (
    view_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    video_id UUID NOT NULL,
    user_id UUID NOT NULL,
    view_date TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (video_id) REFERENCES videos(video_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE for_req (
    data_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    number VARCHAR(9) NOT NULL,
    description TEXT
);