UPDATE tb_projectweb46 SET Title='Test', Content='Test', Author='Yoga Wicaksono', Start_Date='2023/01/05', End_Date='2023/04/10', Techno[1]='JavaScript', Techno[2]='', Techno[3]='', Techno[4]='', Image='image.png' WHERE Id = 8;

SELECT tb_projectweb46.id, title, (End_Date - Start_Date) / 30 as Diff, content, tb_user.name AS author, Techno[1], Techno[2], Techno[3], Techno[4], image FROM tb_projectweb46 LEFT JOIN tb_user ON tb_projectweb46.author = tb_user.id ORDER BY tb_projectweb46.id DESC;

SELECT Id, Title, (End_Date - Start_Date) / 30 as Diff, Content, Author, Techno[1], Techno[2], Techno[3], Techno[4], Image FROM tb_projectweb46;

SELECT tb_projectweb46.Id, Title, End_Date-Start_Date as Diff, TO_CHAR(Start_Date, 'DD-Mon-YYYY') Start_Date, TO_CHAR(End_Date, 'DD-Mon-YYYY') End_Date, Content, Image, tb_user.name AS author, Techno[1], Techno[2], Techno[3], Techno[4] 
FROM tb_projectweb46 LEFT JOIN tb_user ON tb_projectweb46.author = tb_user.id WHERE tb_user.name = 'Yoga';

SELECT Id, Title, End_Date-Start_Date as Diff, TO_CHAR(Start_Date, 'DD-Mon-YYYY') Start_Date, TO_CHAR(End_Date, 'DD-Mon-YYYY') End_Date, Content, Image, Author, Techno[1], Techno[2], Techno[3], Techno[4] FROM tb_projectweb46 WHERE Id = $1;