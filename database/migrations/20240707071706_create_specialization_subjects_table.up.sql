CREATE TABLE specialization_subjects(
    specialization_id CHAR(36) NOT NULL,
    subject_id CHAR(36) NOT NULL,
    PRIMARY KEY (specialization_id, subject_id),
    FOREIGN KEY (specialization_id) REFERENCES specializations(id),
    FOREIGN KEY (subject_id) REFERENCES subjects(id)
) ENGINE=InnoDB;