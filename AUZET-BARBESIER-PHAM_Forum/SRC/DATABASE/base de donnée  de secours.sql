-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : lun. 19 juin 2023 à 17:07
-- Version du serveur : 8.0.31
-- Version de PHP : 8.0.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `test`
--

-- --------------------------------------------------------

--
-- Structure de la table `appreciation_messages`
--

DROP TABLE IF EXISTS `appreciation_messages`;
CREATE TABLE IF NOT EXISTS `appreciation_messages` (
  `id_utilisateurs` tinyint(1) NOT NULL,
  `id_messages` tinyint(1) NOT NULL,
  `etat_like` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `appreciation_messages`
--

INSERT INTO `appreciation_messages` (`id_utilisateurs`, `id_messages`, `etat_like`) VALUES
(23, 2, 1),
(22, 2, 1),
(26, 2, 1),
(22, 3, 1),
(24, 22, 1),
(24, 37, 1),
(24, 38, 1),
(24, 73, 0),
(25, 73, 0),
(24, 70, 1),
(24, 32, 0),
(24, 33, 1),
(50, 62, 1),
(50, 73, 0),
(50, 5, 1),
(50, 46, 0),
(50, 48, 1),
(25, 3, 1),
(25, 31, 1),
(25, 27, 0),
(25, 30, 1),
(25, 47, 0),
(48, 34, 0),
(50, 67, 0),
(50, 57, 1),
(51, 66, 1),
(27, 79, 1);

-- --------------------------------------------------------

--
-- Structure de la table `message`
--

DROP TABLE IF EXISTS `message`;
CREATE TABLE IF NOT EXISTS `message` (
  `id` int NOT NULL AUTO_INCREMENT,
  `dateheure` datetime DEFAULT NULL,
  `contenu` varchar(1024) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `visibilite` tinyint(1) DEFAULT NULL,
  `image` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `update_at` datetime DEFAULT NULL,
  `date_creation` datetime DEFAULT NULL,
  `id_topic` int DEFAULT NULL,
  `id_utilisateurs` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_utilisateurs` (`id_utilisateurs`),
  KEY `id_topic` (`id_topic`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `message`
--

INSERT INTO `message` (`id`, `dateheure`, `contenu`, `visibilite`, `image`, `update_at`, `date_creation`, `id_topic`, `id_utilisateurs`) VALUES
(1, '2023-05-31 15:54:21', 'J\'ai un problème avec mon code, pouvez-vous m\'aider ?', 0, 'image4.jpg', '2023-05-31 15:54:21', NULL, 5, 1),
(2, '2023-05-31 15:54:21', 'Je suis nouveau ici, ravi de vous rejoindre !', 1, 'image5.jpg', '2023-05-31 15:54:21', NULL, 4, 2),
(3, '2023-05-31 15:54:21', 'Bonjour tu a posté ton message en visibilité privé je t\'invite à le réecrire sur un topic public', 0, 'image4.jpg', '2023-05-31 16:54:21', NULL, 5, 5),
(4, '2023-05-31 15:54:21', 'Bonjour Ravi de te connaitre', 0, 'image1.jpg', '2023-05-31 17:54:21', NULL, 4, 1),
(5, '2023-06-10 22:57:00', 'bonjour à tous je me suis toujours demandé quel le meilleur  des deux pour du back', NULL, NULL, NULL, NULL, 3, 1),
(6, '2023-06-10 23:00:00', 'personnellement j\'ai seulement codé que sur du go', NULL, NULL, NULL, NULL, 3, 1),
(7, '2023-06-10 23:02:00', 'franchement je suis pas très disposé à coder sur du js ', NULL, NULL, NULL, NULL, 3, 1),
(8, '2023-06-10 23:05:00', 'y a pas l\'air d\'avoir grand monde en ligne mais  si vous avez des arguments pour / contre n\'hésitez pas', NULL, NULL, NULL, NULL, 3, 1),
(15, '2023-06-10 23:09:00', 'y a pas l\'air d\'avoir grand monde en ligne mais  si vous avez des arguments pour / contre n\'hésitez pas', NULL, NULL, NULL, NULL, 3, 1),
(17, '2023-06-10 23:11:00', 'salut je suis un habitué du back et tant que connaisseur je pense que le go est simple à  utiliser mais que node js est plus flexible', NULL, NULL, NULL, NULL, 3, 2),
(19, '2023-06-10 23:23:00', 'après cela dépends de leur utilisation', NULL, NULL, NULL, NULL, 3, 2),
(20, '2023-06-10 23:25:00', 'si t\'utilise du go back pour faire une base de données qui contiennent des données très sensibles je pense pas que ce soit une bonne idée ...', NULL, NULL, NULL, NULL, 3, 2),
(21, '2023-06-10 23:37:00', 'perso j\'utilise le go que sur des projets qui demande pas trop de sécurité', NULL, NULL, NULL, NULL, 3, 2),
(22, '2023-06-11 18:37:00', 'Bonjour à tous , comme vous le savez python est l\'un des languages les plus connus mais pour autant qu\'est ce qu\'il le rends  autant connu ?  est il plus simple à utiliser que java qui est  LE numéro 1 des langages ? ou simplement à cause des entreprise l', NULL, NULL, NULL, NULL, 6, 2),
(23, '2023-06-11 18:39:00', ' promouvant python ? la question reste toute entière', NULL, NULL, NULL, NULL, 6, 2),
(24, '2023-06-12 08:32:00', 'test', NULL, NULL, NULL, NULL, 3, 2),
(25, '2023-06-12 11:09:00', 'salut  moi je code qu\'en python mais de ce que j\'ai entendu le go c\'est mieux ', NULL, NULL, NULL, NULL, 3, 24),
(26, '2023-06-12 11:18:00', 'node js c\'est vraiment rigide', NULL, NULL, NULL, NULL, 3, 24),
(27, '2023-06-12 11:29:00', 'je préfère fortran', NULL, NULL, NULL, NULL, 3, 27),
(28, '2023-06-12 15:40:00', 'enfin c\'est mon avis', NULL, NULL, NULL, NULL, 3, 2),
(29, '2023-06-12 15:42:00', 'mais je me vois pas passer ma vie à coder en css/js', NULL, NULL, NULL, NULL, 3, 2),
(30, '2023-06-12 15:46:00', 'j\'ai entendu dire que le php est ideal pour la sécurité du back', NULL, NULL, NULL, NULL, 3, 2),
(31, '2023-06-12 20:16:00', 'le python est plus simple c\'est qui le place à la portée du plus grand monde', NULL, NULL, NULL, NULL, 6, 24),
(32, '2023-06-12 21:06:00', 'le ruby m\'a l\'air pas mal ', NULL, NULL, NULL, NULL, 3, 25),
(33, '2023-06-12 21:08:00', 'peut etre mais on parle du go et du js là', NULL, NULL, NULL, NULL, 3, 27),
(34, '2023-06-13 00:24:00', 'meme si le java et plus ancien  il reste moins connus du grand public', NULL, NULL, NULL, NULL, 6, 24),
(35, '2023-06-13 00:25:00', 'tu marque un point', NULL, NULL, NULL, NULL, 6, 27),
(36, '2023-06-13 11:31:00', 'bonjour', NULL, NULL, NULL, NULL, 3, 27),
(37, '2023-06-13 14:36:00', 'ben rip', NULL, NULL, NULL, NULL, 7, 24),
(38, '2023-06-13 14:36:00', 'c\'est triste ', NULL, NULL, NULL, NULL, 7, 24),
(39, '2023-06-13 14:36:00', 'quand memé', NULL, NULL, NULL, NULL, 7, 24),
(40, '2023-06-13 14:36:00', 'croisont les doigts', NULL, NULL, NULL, NULL, 7, 24),
(41, '2023-06-13 14:37:00', 'jsp', NULL, NULL, NULL, NULL, 7, 24),
(42, '2023-06-13 14:38:00', 'rip ', NULL, NULL, NULL, NULL, 7, 24),
(43, '2023-06-13 14:39:00', 'bon ca marche', NULL, NULL, NULL, NULL, 7, 24),
(44, '2023-06-13 14:39:00', 'c\'est cool', NULL, NULL, NULL, NULL, 7, 24),
(45, '2023-06-13 15:32:00', 'message', NULL, NULL, NULL, NULL, 3, 24),
(46, '2023-06-13 15:44:00', 'ce texte est beaucoup trop lonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnng', NULL, NULL, NULL, NULL, 3, 24),
(47, '2023-06-13 15:44:00', 'n\'est ce pas ?', NULL, NULL, NULL, NULL, 3, 24),
(48, '2023-06-13 15:45:00', 'je te rappelle que le flood n\'est pas toléré', NULL, NULL, NULL, NULL, 3, 25),
(52, '2023-06-14 22:54:00', 'salut je suis un habitué du back et tant que connaisseur je pense que le go est simple à  utiliser mais que node js est plus flexibleØje ne suis pas d\'accordØJacquesØ(admin)Ø', NULL, NULL, NULL, NULL, 3, 24),
(53, '2023-06-15 00:22:00', 'le go est plus flexible que node js', NULL, NULL, NULL, NULL, 3, 24),
(54, '2023-06-15 00:23:00', 'je suis pas du tout d\'accord', NULL, NULL, NULL, NULL, 3, 27),
(55, '2023-06-15 00:25:00', 'je te rappelle que le flood n\'est pas toléréØdésoléØeØeØ', NULL, NULL, NULL, NULL, 3, 27),
(56, '2023-06-15 00:28:00', 'le ruby m\'a l\'air pas mal Øje connait pas du tout ce language ØeØeØ', NULL, NULL, NULL, NULL, 3, 27),
(57, '2023-06-16 22:48:00', ' promouvant python ? la question reste toute entièreØet pour le js ?ØJacquesØ(admin)Ø', NULL, NULL, NULL, NULL, 6, 24),
(58, '2023-06-17 01:25:00', 'c\'est subjectif', NULL, NULL, NULL, NULL, 3, 24),
(59, '2023-06-17 01:25:00', 'test', NULL, NULL, NULL, NULL, 3, 24),
(60, '2023-06-17 02:04:00', 'le positionnement des messages devrait marcher', NULL, NULL, NULL, NULL, 6, 24),
(61, '2023-06-17 02:06:00', 'et voila !', NULL, NULL, NULL, NULL, 6, 24),
(62, '2023-06-17 11:55:00', 'je suis content de la nouvelle interface graphique mais il y a de nombreux points à ajouter ', NULL, NULL, NULL, NULL, 8, 24),
(63, '2023-06-17 15:52:00', 'il est vrai que la nouvelle interface est bien plus agréable ', NULL, NULL, NULL, NULL, 9, 24),
(64, '2023-06-17 17:58:00', 'je suis content de la nouvelle interface graphique mais il y a de nombreux points à ajouter Øil manque certaines imagesØzØzØ', NULL, NULL, NULL, NULL, 8, 24),
(65, '2023-06-17 18:48:00', 'il n\'y aurait pas trop de conversation ? ', NULL, NULL, NULL, NULL, 10, 24),
(66, '2023-06-18 13:55:00', 'non je pense que ca va pour le moment ', NULL, NULL, NULL, NULL, 10, 25),
(67, '2023-06-18 13:57:00', 'il n\'y aurait pas trop de conversation ? Øc\'est but d\'un forum d\'avoir pas mal de conversationd\r\nØzØzØ', NULL, NULL, NULL, NULL, 10, 25),
(68, '2023-06-18 16:21:00', 'et voila !Øc\'était compliqué mais ca a marché !ØzØzØ', NULL, NULL, NULL, NULL, 6, 24),
(70, '2023-06-18 20:17:00', 'il reste beaucoup à faire cependant', NULL, NULL, NULL, NULL, 9, 24),
(71, '2023-06-18 20:17:00', 'il reste beaucoup à faire cependantØil faudrait ajouter un système de notificationØzØzØ', NULL, NULL, NULL, NULL, 9, 24),
(73, '2023-06-19 12:07:00', 'j\'aime pas la nouvelle interface', NULL, NULL, NULL, NULL, 9, 49),
(74, '2023-06-19 13:01:00', 'POO ou Programmation dynamique ?', NULL, NULL, NULL, NULL, 3, 25),
(75, '2023-06-19 15:24:00', 'je ne suis pas d\'accordØpourtant les faits sont la ØzØzØ', NULL, NULL, NULL, NULL, 3, 51),
(76, '2023-06-19 15:26:00', 'il manque certaines imagesØj\'ai entendu dire qu\'elles seraient bientôt ajoutéesØzØzØ', NULL, NULL, NULL, NULL, 8, 51),
(77, '2023-06-19 15:27:00', 'il manque certaines imagesØil faut du temps pour faire un forumØzØzØ', NULL, NULL, NULL, NULL, 8, 51),
(78, '2023-06-19 15:40:00', 'j\'ai lu sur un article que c\'était la poule ', NULL, NULL, NULL, NULL, 11, 24),
(79, '2023-06-19 15:45:00', 'j\'ai lu sur un article que c\'était la poule Øil me semble plutot qu\'il était question de l\'oeuf dans l\'articleØzØzØ', NULL, NULL, NULL, NULL, 11, 27),
(88, '2023-06-19 15:47:00', 'j\'ai lu sur un article que c\'était la poule Øun lien avec les dinosaures je crois ØzØzØ', NULL, NULL, NULL, NULL, 11, 24),
(89, '2023-06-19 15:50:00', 'il me semble plutot qu\'il était question de l\'oeuf dans l\'articleØoui tu as raison je suis trompé j\'ai du mal lire l\'article ØgamingØbarbozØ', NULL, NULL, NULL, NULL, 11, 24),
(104, '2023-06-19 15:54:00', 'oui tu as raison je suis trompé j\'ai du mal lire l\'article Ømais pour quelles raisons déja ?ØzØzØ', NULL, NULL, NULL, NULL, 11, 24),
(105, '2023-06-19 15:56:00', 'mais pour quelles raisons déja ?Ø\"En réalité, les poules ne sont pas venus d’œufs de poule, mais d’un autre animal. Des presque-poulets ont un jour pondu un œuf, qui contenait, suite à des mutations génétiques, un gallinacé tel qu’on le connaît\"ØzØzØ', NULL, NULL, NULL, NULL, 11, 27),
(106, '2023-06-19 15:58:00', 'mais pour quelles raisons déja ?Øet voila pourquoiØzØzØ', NULL, NULL, NULL, NULL, 11, 27),
(107, '2023-06-19 16:24:00', 'j\'espère que c\'est plus clair pour toi', NULL, NULL, NULL, NULL, 11, 27);

-- --------------------------------------------------------

--
-- Structure de la table `roles`
--

DROP TABLE IF EXISTS `roles`;
CREATE TABLE IF NOT EXISTS `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `roles`
--

INSERT INTO `roles` (`id`, `nom`) VALUES
(1, 'utilisateur'),
(2, 'administrateur'),
(3, 'moderateur');

-- --------------------------------------------------------

--
-- Structure de la table `topic`
--

DROP TABLE IF EXISTS `topic`;
CREATE TABLE IF NOT EXISTS `topic` (
  `id` int NOT NULL AUTO_INCREMENT,
  `dateheure` datetime DEFAULT NULL,
  `sujet` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `topic`
--

INSERT INTO `topic` (`id`, `dateheure`, `sujet`) VALUES
(3, '2023-05-27 11:04:05', 'go ou node js pour du backend ?'),
(4, '2023-05-26 07:08:54', 'Nouvelles fonctionnalités'),
(5, '2023-05-26 10:08:54', 'Discussions générales'),
(6, '2023-06-11 18:32:00', 'python ou java ? '),
(7, '2023-06-13 11:34:00', 'nouveau topic'),
(8, '2023-06-17 11:54:00', 'le topic de trop ?'),
(9, '2023-06-17 15:27:00', 'nouvelle interface'),
(10, '2023-06-17 18:48:00', 'ca fait pas un peu trop ?'),
(11, '2023-06-19 15:38:00', 'l\'oeuf ou la poule ?');

-- --------------------------------------------------------

--
-- Structure de la table `utilisateurs`
--

DROP TABLE IF EXISTS `utilisateurs`;
CREATE TABLE IF NOT EXISTS `utilisateurs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `prenom` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `mail` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `motdepasse` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `photodeprofil` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `id_roles` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `mail` (`mail`),
  KEY `id_roles` (`id_roles`)
) ENGINE=MyISAM AUTO_INCREMENT=53 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `utilisateurs`
--

INSERT INTO `utilisateurs` (`id`, `nom`, `prenom`, `mail`, `motdepasse`, `photodeprofil`, `id_roles`) VALUES
(1, 'du13', 'Kevin', NULL, 'gqySEzetrdhtf', NULL, 1),
(2, '(admin)', 'Jacques', NULL, 'yugfsetrkjdt', NULL, 3),
(5, 'Thierry', 'Paul', '', 'dzqtgyhfgiuf', NULL, 1),
(6, 'encoding failed', 'gzqrh', 'efshuf@test', 'qefzqeeshtrhdi629865egsrtf', NULL, 1),
(7, 'encoding failed', 'gzqrh', 'alexandre.pham@ynov.com', 'faegyvez6521864dgb5f', NULL, 1),
(8, 'test', 'test', 'test@test', 'test', NULL, 1),
(9, 'alexandre.pham@ynov.com', 'DZGFQYETRHFY', 'DZGTAFEZYHEGTRBTHIU', '65865265', NULL, 1),
(10, 'encoding failed', 'DEZGFYERHSGR', 'DFTEZSETRHDYNTUGYJKT4KAPO35Z4YE5YG', 'D8HYAZFEEFSHGCBILKJLKTA.3GTRDIOBUHNJQEFZYBHJR4UDTRJ', NULL, 1),
(11, 'test', 'xcdbshjfnhtrjk', 'ezytyebhjzuebgfx5644sfsydtg', 'fsgydrt-yrarzhqgthgnez', NULL, 1),
(12, 'fbgyehdgyudr', 'aledbyufgdbgydrj', 'garzehjzgrehjct654iuhfd', 'tftrjgsjgtr', NULL, 1),
(13, 'fctrgcf', 'tyfrgftgh', 'ezsez', '54657', NULL, 1),
(14, 'ab', 'ab', 'ab', 'ab', NULL, 1),
(15, 'a', 'a', 'a', 'a', NULL, 1),
(16, 'qsCdvskjfg', 'ferijtdf', 'dsjfg', 'ezqrt(y', NULL, 1),
(17, 'efzqgtrio', 'sDVfiugh', 'dszerstdtf', 'qsfezsdtrfefzqgtriogfdhtrhjgudgfn', NULL, 1),
(18, 'test', 'qFJEIZETSKRTF', 'qcssibuihgrtdknhfy', 'qdazbezhf', NULL, 1),
(19, 'encoding failed', 'azfteihjr(kltuy;ikuhilj^pkm', 'alexandre.pham@ynov.com0', 'wdskeltr', NULL, 1),
(20, 'test', 'ezqfhyugz(ieul', 'dGfkjegthj,bpomltyeryjufygml', 'azert', NULL, 1),
(21, 'sqcdswftgesijgkl', 'qsFdgjxckhlvnh', 'qDBHUEGRKNHTF IOKLRSJRYOTYGKL.EZOKGETHSRLYGVJBOKL.EZGTRHDIOFK', 'QSfiEGNJqdzfaeijozegtodrkhtfgjobkl', NULL, 1),
(22, 'crypto', 'boomer', 'crytoquimarchepas@notwork.com', '47b375cfaef9bc56f811aa8385f474c6ff3cecbabd7d7a32c23e324bf9970fc1', NULL, 1),
(23, 'sha', '256', '256@sha.com', '51e8ea280b44e16934d4d611901f3d3afc41789840acdff81942c2f65009cd52', NULL, 1),
(24, 'z', 'z', 'z', '594e519ae499312b29433b7dd8a97ff068defcba9755b6d5d00e84c524d67b06', NULL, 1),
(25, 'e', 'e', 'e', '3f79bb7b435b05321651daefd374cdc681dc06faa65e374e38337b88ca046dea', NULL, 1),
(26, 'test', 'ytsggtrd', 'test@jsp.com', 'd74ff0ee8da3b9806b18c877dbf29bbde50b5bd8e4dad7a3a725000feb82e8f1', NULL, 1),
(27, 'barboz', 'gaming', 'gaming@gmail.com', '03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4', NULL, 1),
(28, 'rrr', 'rrrr', 'rrr', '12b0f0dcaefb10c02a83aa9adb025978ddb5512dc04eb39df6811c6a6bf9770c', NULL, 1),
(29, 'zzzzz', 'zzzzz', 'zzzzz', '68a55e5b1e43c67f4ef34065a86c4c583f532ae8e3cda7e36cc79b611802ac07', NULL, 1),
(30, 'essai', 'concluant', 'sucess@mail.com', '1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014', NULL, 1),
(31, 'v', 'v', 'v', '4c94485e0c21ae6c41ce1dfe7b6bfaceea5ab68e40a2476f50208e526f506080', NULL, 1),
(32, 'iii', 'iii', 'iii', 'f5557d4fcf727a981a3c315aca733eefa2996f7c7cdae1fa7e0de28522820bb0', NULL, 1),
(35, '1', '1', '1', '6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b', NULL, 1),
(36, '1', '1', 'mail@', '5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5', NULL, 1),
(37, 'abcd', 'efgh', 'mail@1', '7a5df5ffa0dec2228d90b8d0a0f1b0767b748b0a41314c123075b8289e4e053f', NULL, 1),
(38, 'rip', 'rip', 'rip@gmail.com', '5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5', NULL, 1),
(47, 'essai', 'essai', 'essai', '71b4e190fc7a0aa86f24cb18d88c09bfd8a45292f1ae434fac3c0351f4d838d3', NULL, 1),
(48, 'metro', 'metro', 'metro', '0bc851d8e8861b20da271d23994b9bc2604785313d3369acd3a2f8121b97be61', NULL, 1),
(49, 'hater', 'bashing', 'hater@gmail.com', '41c991eb6a66242c0454191244278183ce58cf4a6bcd372f799e4b9cc01886af', NULL, 1),
(50, 'j', 'j', 'j', '189f40034be7a199f1fa9891668ee3ab6049f82d38c68be70f596eab2e1857b7', NULL, 1),
(44, '__', '__', '__', '9911f4d2b18457c4726664d309385072d295ca69062e99e66250033c13d09441', NULL, 1),
(45, 'è', 'è', 'è', '97c916cd94785d7b9b52ae7013c267854f275691ddaae212bc1c07f639b8a0c7', NULL, 1),
(52, 'clone', 'clone', 'clone', '03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4', NULL, 1),
(51, 'crash', 'crash', 'crash', 'cdb2e0d0f873ce5326e87cf7dec48de8da3043cfc950a7eba05a059150e873f5', NULL, 1);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
