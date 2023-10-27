package parking

type Publisher interface{
	Subscribe(s Subscriber)
}

//TODO: Buat publisher