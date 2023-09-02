# Solution to SQL Murder Mystery

The solution to [SQL Murder Mystery](https://mystery.knightlab.com/) by [Northwestern University Knight Lab](https://knightlab.northwestern.edu/).

## Problem

> A crime has taken place and the detective needs your help. The detective gave you the crime scene report, but you somehow lost it. You vaguely remember that the crime was a ​murder​ that occurred sometime on ​Jan.15, 2018​ and that it took place in ​SQL City​. Start by retrieving the corresponding crime scene report from the police department’s database.

## Schema diagram

![Schema](https://mystery.knightlab.com/schema.png)

## Read the crime scene report

```sql
SELECT description FROM crime_scene_report
WHERE date = 20180115 
    AND type = 'murder' 
    AND city = 'SQL City';
```

> Security footage shows that there were 2 witnesses. The first witness lives at the last house on "Northwestern Dr". The second witness, named Annabel, lives somewhere on "Franklin Ave".

## Read the witness interviews

```sql
WITH witness1 AS (
    SELECT id FROM person
    WHERE address_street_name = 'Northwestern Dr'
    ORDER BY address_number DESC LIMIT 1
), witness2 AS (
    SELECT id FROM person
    WHERE name LIKE '%Annabel%' 
        AND address_street_name = 'Franklin Ave'
), witnesses AS (
    SELECT *, 1 AS witness FROM witness1
    UNION
    SELECT *, 2 AS witness FROM witness2
)
SELECT witness, transcript FROM witnesses
INNER JOIN interview ON witnesses.id = interview.person_id;
```

### Witness 1
> I heard a gunshot and then saw a man run out. He had a "Get Fit Now Gym" bag. The membership number on the bag started with "48Z". Only gold members have those bags. The man got into a car with a plate that included "H42W".
### Witness 2
> I saw the murder happen, and I recognized the killer from my gym when I was working out last week on January the 9th.

## Find the murderer

```sql
WITH gym_members AS (
    SELECT person_id FROM get_fit_now_member 
    INNER JOIN get_fit_now_check_in ON get_fit_now_member.id = get_fit_now_check_in.membership_id
    WHERE id LIKE '48Z%' 
        AND membership_status = 'gold' 
        AND check_in_date = 20180109
)
SELECT person.id, person.name FROM gym_members
INNER JOIN person ON gym_members.person_id = person.id
INNER JOIN drivers_license ON person.license_id = drivers_license.id
WHERE gender = 'male' 
    AND plate_number LIKE '%H42W%';
```

| id    | name          |
|-------|---------------|
| 67318 | Jeremy Bowers |


```sql
INSERT INTO solution VALUES (1, 'Jeremy Bowers');
SELECT value FROM solution;
```

> Congrats, you found the murderer! But wait, there's more... If you think you're up for a challenge, try querying the interview transcript of the murderer to find the real villain behind this crime. If you feel especially confident in your SQL skills, try to complete this final step with no more than 2 queries. Use this same INSERT statement with your new suspect to check your answer.

## Find the real villain behind this crime

```sql
SELECT transcript FROM interview
WHERE person_id = 67318;
```

> I was hired by a woman with a lot of money. I don't know her name but I know she's around 5'5" (65") or 5'7" (67"). She has red hair and she drives a Tesla Model S. I know that she attended the SQL Symphony Concert 3 times in December 2017.

```sql
WITH licenses AS (
    SELECT id FROM drivers_license
    WHERE height BETWEEN 65 AND 67
        AND hair_color = 'red'
        AND gender = 'female'
        AND car_make = 'Tesla'
        AND car_model = 'Model S'
), attenders AS (
    SELECT person_id FROM facebook_event_checkin
    WHERE event_name = 'SQL Symphony Concert'
        AND date BETWEEN 20171200 AND 20171231
    GROUP BY person_id
    HAVING COUNT(person_id) = 3
)
SELECT person.id, person.name FROM attenders
INNER JOIN person ON attenders.person_id = person.id
INNER JOIN licenses ON person.license_id = licenses.id;
```

| id    | name             |
|-------|------------------|
| 99716 | Miranda Priestly |

```sql
INSERT INTO solution VALUES (1, 'Miranda Priestly');
SELECT value FROM solution;
```

> Congrats, you found the brains behind the murder! Everyone in SQL City hails you as the greatest SQL detective of all time. Time to break out the champagne!
