+++
abstract = "The National Center for Biotechnology Information makes available BLAST sequence similarity search (Biotechnology Information (US) 2017a) and health science database search through the Entrez service (Biotechnology Information (US) 2005). In addition to an interactive web interface, BLAST and Entrez provide an application programmer interface to allow programmatic use of these services via the BLAST URL API (Biotechnology Information (US) 2017b) and the Entrez EUtilities (Biotechnology Information (US) 2010). The BioPerl suite (Stajich J.E. 2002) provides access to the BLAST API via Bio::Tools::Run::StandAloneBlastPlus (Fields C. 2017a) and to Entrez via Bio::Tools::EUtilities (Fields C. 2017b). Similarly, Biopython (Cock P.J.A. 2009) provides access via the NCBIWWW function in the Bio.Blast module and functions in Bio.Entrez for EUtilities. Packages within bíogo/ncbi provide Go application programmer interfaces to the NCBI BLAST and EUtilities services. The design of bíogo/ncbi is light weight, allowing the user to make use of the Go language's control structures and data types, rather than imposing a library-specific access approach. In addition to allowing remote BLAST searches, BioPerl and Biopython provide mechanisms to parse XML output from local BLAST search via BioPerl's Bio::SearchIO and Biopythons Bio.Blast NCBIXML. Because of the modular design of bíogo/ncbi, XML encoded output from local BLAST searches or remote downloads can be parsed using the Go standard library's XML package."
abstract_short = ""
authors = ["R Daniel Kortschak", "David L Adelson"]
date = "2017-05-26"
image_preview = ""
math = true
publication_types = ["2"]
publication = "The Journal of Open Source Software"
publication_short = ""
featured = false
title = "bíogo/ncbi: interfaces to NCBI services for the Go language"
url_code = "https://github.com/biogo/ncbi"
url_dataset = ""
url_pdf = "https://github.com/openjournals/joss-papers/raw/master/joss.00234/10.21105.joss.00234.pdf"
url_project = ""
url_slides = ""
url_video = ""

[[links]]
name = "HTML"
url = "http://joss.theoj.org/papers/10.21105/joss.00234"

+++

Biotechnology Information (US), National Center for. 2005. “Entrez Programming Utilities Help.” https://www.ncbi.nlm.nih.gov/books/NBK3836/.

———. 2010. “Entrez Programming Utilities Help.” https://www.ncbi.nlm.nih.gov/books/NBK25501/.

———. 2017a. “Basic Local Alignment Search Tool.” Accessed March 3. https://blast.ncbi.nlm.nih.gov/Blast.cgi.

———. 2017b. “Common Url Api.” Accessed March 3. http://www.ncbi.nlm.nih.gov/BLAST/Doc/urlapi.html.

Cock P.J.A., et al. 2009. “Biopython: Freely Available Python Tools for Computational Molecular Biology and Bioinformatics.” Bioinformatics 25 (11): 1422–3. doi:[10.1093/bioinformatics/btp163](https://doi.org/10.1093/bioinformatics/btp163).

Fields C. 2017a. “Bio-Eutilities.” Accessed March 3. https://github.com/bioperl/bioperl-run.

———. 2017b. “Bio-Eutilities.” Accessed March 3. https://github.com/bioperl/Bio-EUtilities.

Stajich J.E., et al. 2002. “The Bioperl Toolkit: Perl Modules for the Life Sciences.” Genome Research 12 (10): 1611–8. doi:[10.1101/gr.361602](https://doi.org/10.1101/gr.361602).

